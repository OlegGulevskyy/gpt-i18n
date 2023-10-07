package main

import (
	"fmt"
	"os"
	"strings"
)

// All github actions when triggered prefix env variables with this value
const GH_ACTIONS_INPUT_PREF = "INPUT_"

// Use these "values" to identify Env variables related to gpt-i18n
const GPT_I18N_GLOB_VAR = "GPT_GLOB_"
const GPT_I18N_TRANSLATE_TO_LANGS_VAR = "GPT_TRANSLATE_TO_LANGS"
const INPUT_GPT_OPENAI_KEY = "GPT_OPENAI_KEY"

func getGhActionEnv(env string) string {
	return fmt.Sprintf("%s%s", GH_ACTIONS_INPUT_PREF, env)
}

/**
* Manual handling of env variables
* because the goal is to provide multiple env variables
* but sticking to a common naming convention.
 */
func getEnvValue(env string) string {
	return strings.Split(env, "=")[1]
}

type InputValues struct {
	/** Directories that run gpt-i18n against. */
	dirs []string
	/** Translate files / differences to languages, comma separated. */
	targetLangs []string
	/** OpenAI API Key to use. */
	openAiApiKey string
}

func getGhInputVariables() InputValues {
	envs := os.Environ()

	dirKey := getGhActionEnv(GPT_I18N_GLOB_VAR)
	targetLangsKey := getGhActionEnv(GPT_I18N_TRANSLATE_TO_LANGS_VAR)
	openAiEnvKey := getGhActionEnv(INPUT_GPT_OPENAI_KEY)

	dirs := []string{}
	targetLangs := []string{}
	openAiApi := ""

	inputValues := InputValues{
		dirs,
		targetLangs,
		openAiApi,
	}

	for _, env := range envs {
		if strings.Contains(env, dirKey) {
			inputValues.dirs = append(inputValues.dirs, getEnvValue(env))
		}

		if strings.Contains(env, targetLangsKey) {
			envValue := getEnvValue(env)
			langsSplit := strings.Split(envValue, ",")
			inputValues.targetLangs = append(targetLangs, langsSplit...)
		}

		if strings.Contains(env, openAiEnvKey) {
			inputValues.openAiApiKey = getEnvValue(env)
		}
	}

	return inputValues
}

func ensureRequiredEnvsPresent(envValues *InputValues) {
	if envValues.dirs == nil || len(envValues.dirs) == 0 {
		ErrLog(fmt.Sprintf("No directories to run gpt-i18n in. Set %s*ANYTHING* in 'with' value", GPT_I18N_GLOB_VAR))
		os.Exit(1)
	}

	if envValues.openAiApiKey == "" {
		ErrLog(fmt.Sprintf("No OpenAI API Key provided. Set %s to your key in 'with' value", INPUT_GPT_OPENAI_KEY))
		os.Exit(1)
	}

	if envValues.targetLangs == nil || len(envValues.targetLangs) == 0 {
		ErrLog(fmt.Sprintf("No target languages provided. Set %s to your key in 'with' value separated by comma. Example: 'ru,fr' ", GPT_I18N_TRANSLATE_TO_LANGS_VAR))
		os.Exit(1)
	}
}

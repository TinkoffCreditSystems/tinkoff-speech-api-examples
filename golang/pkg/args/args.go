package args

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

type RecognizeOptions struct {
	InputFile                   *os.File
	Encoding                    *string
	Rate                        *int
	NumChannels                 *int
	MaxAlternatives             *int
	DoNotPerformVad             *bool
	SilenceDurationThreshold    *float64
	LanguageCode                *string
	DisableAutomaticPunctuation *bool
	DisableProfanityFilter      *bool
	EnableDenormalization       *bool
}

type StreamingRecognizeOptions struct {
	RecognizeOptions
	SingleUtterance *bool
	InterimResults  *bool
}

type StreamingSynthesizeOptions struct {
	InputText  *string
	OutputFile *os.File
	Encoding   *string
	Rate       *int
}

func addRecognizeOptions(parser *argparse.Parser) *RecognizeOptions {
	inputFile := parser.File("i", "input", os.O_RDONLY, 0400, &argparse.Options{
		Required: true,
		Help:     "Input audio file to recognize.",
	})

	encoding := parser.Selector("e", "encoding", []string{
		"LINEAR16",
		"MPEG_AUDIO",
		"ALAW",
		"MULAW",
	}, &argparse.Options{
		Required: true,
		Help:     "Audio encoding.",
	})

	rate := parser.Int("r", "rate", &argparse.Options{
		Required: true,
		Help:     "Audio sampling rate.",
	})

	numChannels := parser.Int("c", "num-channels", &argparse.Options{
		Required: true,
		Help:     "Number of audio channels.",
	})

	maxAlternatives := parser.Int("", "max-alternatives", &argparse.Options{
		Help:    "Number of speech recognition alternatives to return.",
		Default: 1,
	})

	doNotPerformVad := parser.Flag("", "do-not-perform-vad", &argparse.Options{
		Help: "Specify this to disable voice activity detection. All audio is processed " +
			"as though it were a single utterance.",
		Default: false,
	})

	silenceDurationThreshold := parser.Float("", "silence-duration-threshold", &argparse.Options{
		Help: "Silence threshold in seconds for VAD to assume the current utterance is ended and " +
			"the next utterance shall begin.",
		Default: 0.6,
	})

	languageCode := parser.String("", "language-code", &argparse.Options{
		Help:    "Language for speech recognition.",
		Default: "ru-RU",
	})

	disableAutomaticPunctuation := parser.Flag("", "disable-automatic-punctuation", &argparse.Options{
		Help:    "Specify this to disable automatic punctuation in recognition results.",
		Default: false,
	})

	disableProfanityFilter := parser.Flag("", "disable-profanity-filter", &argparse.Options{
		Help:    "Specify this to disable profanity filter.",
		Default: false,
	})

	enableDenormalization := parser.Flag("", "enable-denormalization", &argparse.Options{
		Help:    "Enables automatic conversion of numerals from text to numeric form",
		Default: false,
	})

	return &RecognizeOptions{
		InputFile:                   inputFile,
		Encoding:                    encoding,
		Rate:                        rate,
		NumChannels:                 numChannels,
		MaxAlternatives:             maxAlternatives,
		DoNotPerformVad:             doNotPerformVad,
		SilenceDurationThreshold:    silenceDurationThreshold,
		LanguageCode:                languageCode,
		DisableAutomaticPunctuation: disableAutomaticPunctuation,
		DisableProfanityFilter:      disableProfanityFilter,
		EnableDenormalization:       enableDenormalization,
	}
}

func addStreamingRecognizeOptions(parser *argparse.Parser) *StreamingRecognizeOptions {
	recognizeOptions := addRecognizeOptions(parser)
	singleUtterance := parser.Flag("", "single-utterance", &argparse.Options{
		Help:    "Single utterance",
		Default: false,
	})
	interimResults := parser.Flag("", "interim-results", &argparse.Options{
		Help:    "Whether to enable interim results",
		Default: false,
	})
	return &StreamingRecognizeOptions{
		RecognizeOptions: *recognizeOptions,
		SingleUtterance:  singleUtterance,
		InterimResults:   interimResults,
	}
}

func addStreamingSynthesizeOptions(parser *argparse.Parser) *StreamingSynthesizeOptions {
	inputText := parser.String("i", "input-text", &argparse.Options{
		Required: true,
		Help:     "Input text to synthesize.",
	})
	outputFile := parser.File("o", "output", os.O_CREATE|os.O_WRONLY, 0664, &argparse.Options{
		Required: true,
		Help:     "Output wav to save.",
	})
	encoding := parser.Selector("e", "encoding", []string{"RAW_OPUS", "LINEAR16"}, &argparse.Options{
		Required: true,
		Help:     "Audio encoding.",
	})
	rate := parser.Int("r", "rate", &argparse.Options{
		Required: true,
		Help:     "Audio sample rate.",
	})

	return &StreamingSynthesizeOptions{
		InputText:  inputText,
		OutputFile: outputFile,
		Encoding:   encoding,
		Rate:       rate,
	}
}

func ParseRecognizeOptions() *RecognizeOptions {
	parser := argparse.NewParser("recognize", "Recognize file in a non-streaming manner")
	recognizeOptions := addRecognizeOptions(parser)
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return nil
	}
	return recognizeOptions
}

func ParseStreamingRecognizeOptions() *StreamingRecognizeOptions {
	parser := argparse.NewParser("streaming_recognize", "Recognize file in a streaming manner")
	streamingRecognizeOptions := addStreamingRecognizeOptions(parser)
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return nil
	}
	return streamingRecognizeOptions
}

func ParseStreamingSynthesizeOptions() *StreamingSynthesizeOptions {
	parser := argparse.NewParser("streaming_synthesize", "Synthesize a wav file from text")
	streamingSynthesizeOptions := addStreamingSynthesizeOptions(parser)
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf(parser.Usage(err))
		return nil
	}
	return streamingSynthesizeOptions
}

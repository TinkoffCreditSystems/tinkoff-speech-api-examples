# Golang NodeJS examples

You will need go with modules support and opus to run this examples.

## Usage

### Install the requirements

On MacOS:

```
$ brew install opus opusfile
```

On Ubuntu:

```
$ apt install libopus-dev libopusfile-dev
```

### Build example binaries

Build binaries:

```
$ go build cmd/recognize/recognize.go
$ go build cmd/recognize_stream/recognize_stream.go
$ go build cmd/synthesize_stream/synthesize_stream.go
```

### Basic recognition examples

Run basic (non-streaming) speech recognition example:

```
$ ./recognize -e MPEG_AUDIO -r 16000 -c 1 -i ../audio/sample_1.mp3
```

To disable automatic punctuation and get up to 3 recognition alternatives:

```
$ ./recognize -e MPEG_AUDIO -r 16000 -c 1 --disable-automatic-punctuation --max-alternatives 3 -i ../audio/sample_1.mp3
```

To disable profanity filter:

```
$ ./recognize -e MPEG_AUDIO -r 16000 -c 1 --disable-profanity-filter -i ../audio/sample_5.mp3
```

To enable conversion numerals from text to numeric form:

```
$ ./recognize -e MPEG_AUDIO -r 16000 -c 1 --enable-denormalization -i ../audio/sample_6.mp3
```

Run streaming speech recognition with interim results:

```
$ ./recognize_stream -e MPEG_AUDIO -r 16000 -c 1 --interim-results -i ../audio/sample_1.mp3
```

Specify longer silence timeout for voice activity detection (you will probably need longer audio to actually see the difference):

```
$ ./recognize_stream -e MPEG_AUDIO -r 16000 -c 1 --interim-results --silence-duration-threshold 1.2 -i ../audio/sample_1.mp3
```

Return just the first recognized utterance and halt (you will probably need longer audio to actually see the difference):

```
$ ./recognize_stream -e MPEG_AUDIO -r 16000 -c 1 --interim-results --single-utterance -i ../audio/sample_1.mp3
```

### Basic synthesis examples

To run basic speech synthesis and save result to wav:

```
$ ./synthesize_stream -r 48000 -e LINEAR16 -i "И мысли тоже тяжелые и медлительные, падают неторопливо и редко одна за другой, точно песчинки в разленившихся песочных часах." -o output_1.wav
```

You can also specify alternative stress with `0` sign after a vowel:

```
$ ./synthesize_stream -r 48000 -e LINEAR16 -i "За0мок - замо0к." -o output_2.wav
```

Feel free to use arabic numerals and named entities:

```
$ ./synthesize_stream -r 48000 -e LINEAR16 -i "Газета Times, 03 января 2009 года - Канцлер на грани ради второго спасения банков." -o output_3.wav
```

For now, `LINEAR16` does not support samples rates other than 48kHz. Use `RAW_OPUS` (`node-opus` required) to specify different sample rates:

```
$ ./synthesize_stream -r 16000 -e RAW_OPUS -i "Привет, мир." -o output_4.wav
```

## Generate Protobuf and gRPC definitions (optional)

In case of API changes (`*.proto` files in `apis` directory),
you may regenerate Protobuf and gRPC definitions by simply running the following script
(no need to re-clone the whole repo):

```
$ ./sh/generate_protobuf.sh
```

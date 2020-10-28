// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes. SSML input
// must be valid, well-formed SSML. Some alphabets might not be available with all
// the voices (for example, Cyrillic might not be read at all by English voices)
// unless phoneme mapping is used. For more information, see How it Works
// (https://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html).
func (c *Client) SynthesizeSpeech(ctx context.Context, params *SynthesizeSpeechInput, optFns ...func(*Options)) (*SynthesizeSpeechOutput, error) {
	if params == nil {
		params = &SynthesizeSpeechInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SynthesizeSpeech", params, optFns, addOperationSynthesizeSpeechMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SynthesizeSpeechOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SynthesizeSpeechInput struct {

	// The format in which the returned output will be encoded. For audio stream, this
	// will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json. When pcm
	// is used, the content returned is audio/pcm in a signed 16-bit, 1 channel (mono),
	// little-endian format.
	//
	// This member is required.
	OutputFormat types.OutputFormat

	// Input text to synthesize. If you specify ssml as the TextType, follow the SSML
	// format for the input text.
	//
	// This member is required.
	Text *string

	// Voice ID to use for the synthesis. You can get a list of available voice IDs by
	// calling the DescribeVoices
	// (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html) operation.
	//
	// This member is required.
	VoiceId types.VoiceId

	// Specifies the engine (standard or neural) for Amazon Polly to use when
	// processing input text for speech synthesis. Using a voice that is not supported
	// for the engine selected will result in an error.
	Engine types.Engine

	// Optional language code for the Synthesize Speech request. This is only necessary
	// if using a bilingual voice, such as Aditi, which can be used for either Indian
	// English (en-IN) or Hindi (hi-IN). If a bilingual voice is used and no language
	// code is specified, Amazon Polly will use the default language of the bilingual
	// voice. The default language for any voice is the one returned by the
	// DescribeVoices
	// (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html) operation
	// for the LanguageCode parameter. For example, if no language code is specified,
	// Aditi will use Indian English rather than Hindi.
	LanguageCode types.LanguageCode

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon is
	// the same as the language of the voice. For information about storing lexicons,
	// see PutLexicon
	// (https://docs.aws.amazon.com/polly/latest/dg/API_PutLexicon.html).
	LexiconNames []*string

	// The audio frequency specified in Hz. The valid values for mp3 and ogg_vorbis are
	// "8000", "16000", "22050", and "24000". The default value for standard voices is
	// "22050". The default value for neural voices is "24000". Valid values for pcm
	// are "8000" and "16000" The default value is "16000".
	SampleRate *string

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []types.SpeechMarkType

	// Specifies whether the input text is plain text or SSML. The default value is
	// plain text. For more information, see Using SSML
	// (https://docs.aws.amazon.com/polly/latest/dg/ssml.html).
	TextType types.TextType
}

type SynthesizeSpeechOutput struct {

	// Stream containing the synthesized speech.
	AudioStream io.ReadCloser

	// Specifies the type audio stream. This should reflect the OutputFormat parameter
	// in your request.
	//
	//     * If you request mp3 as the OutputFormat, the ContentType
	// returned is audio/mpeg.
	//
	//     * If you request ogg_vorbis as the OutputFormat,
	// the ContentType returned is audio/ogg.
	//
	//     * If you request pcm as the
	// OutputFormat, the ContentType returned is audio/pcm in a signed 16-bit, 1
	// channel (mono), little-endian format.
	//
	//     * If you request json as the
	// OutputFormat, the ContentType returned is audio/json.
	ContentType *string

	// Number of characters synthesized.
	RequestCharacters *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSynthesizeSpeechMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpSynthesizeSpeech{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpSynthesizeSpeech{})
	if err != nil {
		return err
	}
	if err := awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err := addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err := v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err := addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err := addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err := awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err := addClientUserAgent(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err := addOpSynthesizeSpeechValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opSynthesizeSpeech(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSynthesizeSpeech(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "SynthesizeSpeech",
	}
}

package ocr

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Vision -framework CoreGraphics -framework AppKit

#import <Foundation/Foundation.h>
#import <Vision/Vision.h>
#import <CoreGraphics/CoreGraphics.h>
#import <AppKit/AppKit.h>

char* performOCR(const char* imagePath) {
    @autoreleasepool {
        NSString *path = [NSString stringWithUTF8String:imagePath];
        NSImage *image = [[NSImage alloc] initWithContentsOfFile:path];

        if (!image) {
            return strdup("Error: Failed to load image");
        }

        // Convert NSImage to CGImage
        CGImageSourceRef source = CGImageSourceCreateWithData((CFDataRef)[image TIFFRepresentation], NULL);
        CGImageRef cgImage = CGImageSourceCreateImageAtIndex(source, 0, NULL);
        CFRelease(source);

        if (!cgImage) {
            return strdup("Error: Failed to convert image");
        }

        // Create Vision request handler
        VNImageRequestHandler *handler = [[VNImageRequestHandler alloc] initWithCGImage:cgImage options:@{}];

        // Create text recognition request
        VNRecognizeTextRequest *request = [[VNRecognizeTextRequest alloc] init];
        request.recognitionLevel = VNRequestTextRecognitionLevelAccurate;
        // Only use Simplified Chinese for better accuracy
        request.recognitionLanguages = @[@"zh-Hans"];
        request.usesLanguageCorrection = YES;

        NSError *error = nil;
        [handler performRequests:@[request] error:&error];

        CGImageRelease(cgImage);

        if (error) {
            NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", error.localizedDescription];
            return strdup([errorMsg UTF8String]);
        }

        // Extract recognized text
        NSMutableString *recognizedText = [NSMutableString string];
        NSArray<VNRecognizedTextObservation *> *observations = request.results;

        for (VNRecognizedTextObservation *observation in observations) {
            VNRecognizedText *topCandidate = [observation topCandidates:1].firstObject;
            if (topCandidate) {
                [recognizedText appendString:topCandidate.string];
                [recognizedText appendString:@"\n"];
            }
        }

        if (recognizedText.length == 0) {
            return strdup("No text recognized");
        }

        return strdup([recognizedText UTF8String]);
    }
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

// RecognizeText performs OCR on the image at the given path
func RecognizeText(imagePath string) (string, error) {
	cPath := C.CString(imagePath)
	defer C.free(unsafe.Pointer(cPath))

	result := C.performOCR(cPath)
	defer C.free(unsafe.Pointer(result))

	text := C.GoString(result)

	if len(text) > 6 && text[:6] == "Error:" {
		return "", errors.New(text[7:])
	}

	if text == "No text recognized" {
		return "", errors.New("no text found in image")
	}

	return text, nil
}

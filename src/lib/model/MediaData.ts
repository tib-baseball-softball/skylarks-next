export interface MediaData {
    publicUrl: string;
    properties: {
        title: string | null;
        alternative: string | null;
        description: string | null;
        link: string | null;
        linkData: any | null; // You might want to replace 'any' with the appropriate type for your link data
        mimeType: string;
        type: string;
        filename: string;
        originalUrl: string;
        uidLocal: number;
        fileReferenceUid: number;
        size: string;
        dimensions: {
            width: number;
            height: number;
        };
        cropDimensions: {
            width: number;
            height: number;
        };
        crop: {
            default?: {
                cropArea: {
                    x: number;
                    y: number;
                    width: number;
                    height: number;
                };
                selectedRatio: string;
                focusArea: any | null; // You might want to replace 'any' with the appropriate type for your focus area data
            };
            social?: {
                cropArea: {
                    x: number;
                    y: number;
                    width: number;
                    height: number;
                };
                selectedRatio: string;
                focusArea: any | null; // You might want to replace 'any' with the appropriate type for your focus area data
            };
        };
        autoplay: boolean | null;
        extension: string;
    }
}
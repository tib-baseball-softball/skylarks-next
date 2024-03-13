import type {Typo3DataObject} from "$lib/model/Typo3DataObject";

export interface Typo3ContentObject extends Typo3DataObject {
    type: string;
    colPos: number;
    categories: string;
    appearance: {
        layout: string;
        frameClass: string;
        spaceBefore: string;
        spaceAfter: string;
    };
    content: {
        header: string;
        subheader: string;
        headerLayout: number;
        headerPosition: string;
        headerLink: string;
        bodytext: string;
        enlargeImageOnClick?: boolean;
        gallery?: {
            position: {
                horizontal: string;
                vertical: string;
                noWrap: boolean;
            };
            width: number;
            count: {
                files: number;
                columns: number;
                rows: number;
            };
            columnSpacing: number;
            border: {
                enabled: boolean;
                width: number;
                padding: number;
            };
            rows: any[]; // You might want to replace 'any' with the appropriate type for your gallery data
        };
    };
}
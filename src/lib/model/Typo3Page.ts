import type {Typo3DataObject} from "$lib/model/Typo3DataObject";
import type {MediaData} from "$lib/model/MediaData";
import type {Typo3ContentObject} from "$lib/model/Typo3ContentObject";
import type {Breadcrumb} from "$lib/model/Breadcrumb";

export interface PageMetaData {
    title: string;
    subtitle: string;
    abstract: string;
    description: string;
    keywords: string;
    canonical: string;
    robots: {
        noIndex: boolean;
        noFollow: boolean;
    };
    author: string;
    authorEmail: string;
    ogTitle: string;
    ogDescription: string;
    ogImage: null | MediaData;
    twitterTitle: string;
    twitterDescription: string;
    twitterCard: string;
    twitterImage: null | MediaData;
}

export interface PageAppearanceData {
    layout: string;
    backendLayout: string;
}

export interface Typo3Page extends Typo3DataObject {
    id: number;
    type: string;
    slug: string;
    media: MediaData[];
    meta: PageMetaData;
    categories: string;
    breadcrumbs: Breadcrumb[];
    appearance: PageAppearanceData;
    content: {
        colPos0?: Typo3ContentObject[];
        colPos1?: Typo3ContentObject[];
    };
}

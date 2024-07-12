import type {Media} from "$lib/model/MediaData";
import type {BaseballPosition} from "$lib/types/BaseballPosition";
import type {Handedness} from "$lib/types/Handedness";

export interface Player {
    uid: number;
    pid: number;
    fullname: string;
    firstname: string;
    lastname: string;
    birthday: number;
    admission: string;
    number: string;
    throwing: Handedness;
    batting: Handedness;
    coach: string;
    slug: string;
    bsm_id: number;
    media: Media[];
    positions: BaseballPosition[];
    teams: {uid: number, name: string}[]
}
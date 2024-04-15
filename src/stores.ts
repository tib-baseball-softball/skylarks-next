import type {Writable} from "svelte/store";
import {localStorageStore} from "@skeletonlabs/skeleton";

const currentYear = new Date().getFullYear()

export const selectedSeason: Writable<number> = localStorageStore('selectedSeason', currentYear)
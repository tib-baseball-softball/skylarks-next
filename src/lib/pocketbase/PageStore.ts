import { type Readable } from "svelte/store"
import { type ListResult } from "pocketbase"

export interface PageStore<T = any> extends Readable<ListResult<T>> {
  setPage(newpage: number): Promise<void>

  next(): Promise<void>

  prev(): Promise<void>
}

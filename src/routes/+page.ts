import { writable, readable, derived } from "svelte/store";

export const _data = writable([])

export let _read_data = derived(_data, ($_data) => {
    if($_data) {
        return $_data      
    }
    return []
})

export function _newTask(taskDesc: string) {
    if (!taskDesc) return null
    fetch("http://127.0.0.1:5050/new/" + taskDesc)
}
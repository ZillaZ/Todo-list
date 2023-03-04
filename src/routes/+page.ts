import { writable, derived } from "svelte/store";

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
    _updateStores()
}

export function _updateStores() {
    fetch("http://127.0.0.1:5050")
    .then(response => response.json()
    .then(_read_data => {
        _data.set(_read_data)
    })).catch(error => {
        console.log(error)
        return []
    })
}

export function _deleteTask(taskDesc: string) {
    fetch("http://127.0.0.1:5050/del/" + taskDesc)
    _updateStores()
}

export function _updateStatus(taskDesc: string) {
    fetch("http://127.0.0.1:5050/" + taskDesc)
    _updateStores()
}
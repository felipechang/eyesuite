import {writable} from "svelte/store";

export const errorStore = writable({
    title: "",
    error: ""
});
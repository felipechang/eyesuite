import {writable} from "svelte/store";
import {getServer, postServer} from "$lib/stores/wrapper";

const init: IProfileInit[] = [];

const ENDPOINT = "/api/profiles";

function createStore() {
    const {subscribe, set} = writable(init);
    return {
        findOne: (profiles: IProfileInit[], name: string): IProfileInit => {
            const p = profiles.find((e: IProfileInit) => {
                return e.name === name;
            });
            if (!p) {
                return {name: "", plugin: "", template: "", mapping: []}
            }
            return p;
        },
        init: (): IProfileInit[] => init,
        mount: async () => {
            const response = await getServer(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (profiles: IProfileInit[]) => {
            const response = await postServer(ENDPOINT, JSON.stringify(profiles));
            const data = await response.json();
            set(data.data);
        },
    };
}

export const profileListStore = createStore();
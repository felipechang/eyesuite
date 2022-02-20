import {postServer} from "$lib/stores/wrapper";

function createStore() {
    return {
        submit: async (template: string, files: File) => {
            const formData = new FormData()
            formData.append("file", files);
            formData.append("template", template);
            const response = await postServer(`api/postImage`, formData);

            // invalid credentials
            if (response.status === 400) {
                return {
                    error: true,
                    msg: "Invalid NetSuite credentials. Please contact your administrator."
                };
            }
            const res = await response.json();
            return {error: false, msg: res.data};
        },
        read: async (link: string, files: File) => {
            const formData = new FormData()
            formData.append("file", files);
            const response = await postServer(link, formData);
            const data = await response.json();
            return data.data;
        }
    };
}

export const readerStore = createStore();
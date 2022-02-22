import { makeHeaders } from "$lib/stores/wrapper";
import { loginStore } from "$lib/stores/login";
const submit = async (template, files) => {
    const formData = new FormData();
    formData.append("file", files);
    formData.append("template", template);
    const headers = makeHeaders();
    const response = await fetch("/api/postImage", {
        headers: headers,
        method: 'POST',
        body: formData
    });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh();
        return submit(template, files);
    }
    // invalid credentials
    if (response.status === 400) {
        return {
            error: true,
            msg: "Invalid NetSuite credentials. Please contact your administrator."
        };
    }
    return { error: false, msg: await response.json() };
};
const read = async (link, file) => {
    const formData = new FormData();
    formData.append("file", file);
    const headers = makeHeaders();
    const response = await fetch(link, {
        headers: headers,
        method: 'POST',
        body: formData
    });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh();
        return read(link, file);
    }
    // invalid credentials
    if (response.status === 400) {
        return {
            error: true,
            msg: "Invalid NetSuite credentials. Please contact your administrator."
        };
    }
    return { error: false, msg: await response.json() };
};
function createStore() {
    return {
        submit,
        read,
    };
}
export const readerStore = createStore();
//# sourceMappingURL=reader.js.map
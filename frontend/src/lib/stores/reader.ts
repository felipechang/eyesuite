import {makeHeaders} from "$lib/stores/wrapper";
import {loginStore} from "$lib/stores/login";

interface ISubmit {
    error: boolean,
    msg: string
}

const read = async (link: string, file: File, fields: { name: string, value: string }[]): Promise<ISubmit> => {
    const formData = new FormData()
    formData.append("file", file);
    for (let i = 0; i < fields.length; i++) {
        formData.append(fields[i].name, fields[i].value);
    }
    const headers = makeHeaders();
    const response = await fetch(link, {
        headers: headers,
        method: 'POST',
        body: formData,
    });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh();
        return read(link, file, fields);
    }
    // invalid credentials
    if (response.status === 400) {
        return {
            error: true,
            msg: "Invalid NetSuite credentials. Please contact your administrator."
        };
    }
    const evalResponse = await response.clone();
    const evalMsg = await evalResponse.json();
    if (evalMsg === "NotFoundException" || evalMsg === "FormatException") {
        return {
            error: true,
            msg: "Incorrect barcode type"
        };
    }
    return {error: false, msg: await response.json()};
}

const submit = async (template: string, files: File): Promise<ISubmit> => {
    const formData = new FormData()
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
    return {error: false, msg: await response.json()};
}

function createStore() {
    return {
        submit,
        read,
    };
}

export const readerStore = createStore();
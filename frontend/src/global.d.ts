/// <reference types="@sveltejs/kit" />

interface IMapping {
    labelText: string
    name: string
    value: string
    placeholder: string
    readonly: boolean
    hidden: boolean
}

interface IConnectionInit {
    version: string
    ws_url: string
    realm: string
    consumer_key: string
    consumer_secret: string
    token_id: string
    token_secret: string
}

interface IPluginInit {
    name: string
    description: string
    enabled: boolean
    mapping: IMapping[]
}

interface IProfileInit {
    name: string
    template: string
    plugin: string
    mapping: IMapping[]
}

interface IReaderInit {
    profiles: IProfileInit[]
    plugins: IPluginInit[]
}

interface IUserInit {
    name: string
    username: string
    key: string
    enabled: boolean
    admin: boolean
}
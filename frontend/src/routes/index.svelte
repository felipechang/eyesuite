{#if connectionStore.validate(connection)}
    <Grid>
        <SubNav
                profiles="{profiles}"
                loadMappings="{loadMappings}"
                updateSelectChange="{updateSelectChange}"
        />
        {#if profileSelected}
            <Camera
                    success={success}
                    bind:files="{files}"
            />
            {#if files}
                <Form
                        mappings={mappings}
                        pluginLink="{pluginLink}"
                        pluginField="{pluginField}"
                        template="{template}"
                        files="{files}"
                        bind:success="{success}"
                        loadMappings="{loadMappings}"
                        updateSelectChange="{updateSelectChange}"
                />
            {/if}
        {/if}
    </Grid>
{:else}
    <p>Please complete NetSuite connection settings</p>
{/if}

<script lang="ts">
    import {connectionStore} from "$lib/stores/connection";
    import {Grid} from "carbon-components-svelte";
    import Camera from "./reader/Camera.svelte";
    import SubNav from "./reader/SubNav.svelte";
    import Form from "./reader/Form.svelte";
    import {onMount} from "svelte";
    import {pluginListStore} from "$lib/stores/plugins";
    import {profileListStore} from "$lib/stores/profiles";

    let profileSelected = "";
    let pluginLink = "";
    let pluginField = "";
    let template = "";
    let files = "";
    let success = "";

    let mappings = [];

    let connection = connectionStore.init();
    let profiles = profileListStore.init();
    let plugins = pluginListStore.init();

    const loadMappings = async () => {
        plugins = await pluginListStore.mount();
        profiles = (await profileListStore.mount()).filter((p) => {
            // filter disabled plugins
            for (let i = 0; i < plugins.length; i++) {
                if (p.plugin === plugins[i].name) {
                    return plugins[i].enabled;
                }
            }
            return false;
        });
    }

    /** updateSelectChange update reader selection */
    const updateSelectChange = (profileName) => {
        if (!profileName) {
            profileSelected = "";
            pluginLink = "";
            pluginField = "";
            mappings = [];
            return;
        }
        profileSelected = profileName;
        const e = profileListStore.findOne(profiles, profileName);
        if (!e.plugin) {
            return;
        }
        const f = pluginListStore.findOne(plugins, e.plugin);
        mappings = e.mapping.concat(f.enabled ? f.mapping : []);
        pluginLink = f.endpoint;
        template = e.template;
        pluginField = f.endpointField;
    }

    onMount(async () => {
        connection = await connectionStore.mount()
    });
</script>
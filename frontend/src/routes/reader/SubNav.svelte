<Row>
    <Column>
        <Title
                details="Take a photo and fill in the details"
                name="Reader"
        />
    </Column>
    <Column>
        <Select
                bind:selected={current}
                labelText="Select Profile:"
                on:input={handleSelectChange}
        >
            <SelectItem text="" value=""/>
            {#each profiles as profile}
                <SelectItem text="{profile.name}" value="{profile.name}"/>
            {/each}
        </Select>
    </Column>
</Row>

<script lang="ts">
    import {Column, Row, Select, SelectItem} from "carbon-components-svelte";
    import {profileListStore} from "$lib/stores/profiles";
    import {pluginListStore} from "$lib/stores/plugins";
    import Title from "$lib/components/Title.svelte";
    import {onMount} from "svelte";

    export let profileSelected;
    export let pluginLink;
    export let pluginField;
    export let template;
    export let mappings;

    let current;
    let profiles = profileListStore.init();
    let plugins = pluginListStore.init();

    onMount(async () => {
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

        // set last option
        current = localStorage.getItem("profile-selected");
        updateSelectChange(current);
    });

    /** handleSelectChange export selected profile name */
    const handleSelectChange = (e) => {
        if (!e.target.value) {
            profileSelected = "";
            pluginLink = "";
            pluginField = "";
            mappings = [];
            return;
        }
        localStorage.setItem("profile-selected", e.target.value);
        updateSelectChange(e.target.value);
    }

    /** updateSelectChange update reader selection */
    const updateSelectChange = (profileName) => {
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
</script>
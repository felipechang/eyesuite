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
    import Title from "$lib/components/Title.svelte";
    import {onMount} from "svelte";
    import {pluginListStore} from "$lib/stores/plugins";

    export let profileSelected;
    export let mappings;

    let profiles = [];
    let plugins = [];
    let current;

    onMount(async () => {

        if (profiles.length !== 0) {
            return;
        }

        // filter disabled plugins
        plugins = await pluginListStore.mount()
        let profileList = await profileListStore.mount()
        profiles = profileList.filter((p) => {
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
        profileSelected = e.target.value;
        localStorage.setItem("profile-selected", profileSelected);
        updateSelectChange(profileSelected);
    }

    /** updateSelectChange update reader selection */
    const updateSelectChange = (p) => {
        profileSelected = p;
        const e = profileListStore.findOne(profiles, p);
        if (!e.plugin) {
            return;
        }
        const f = pluginListStore.findOne(plugins, e.plugin);
        mappings = e.mapping.concat(f.enabled ? f.mapping : []);
    }
</script>
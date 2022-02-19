<Select bind:selected="{profile.plugin}" on:input={handleSelectChange}>
    <SelectItem text="" value=""/>
    {#each plugins as plugin}
        {#if plugin.enabled}
            <SelectItem text="{plugin.name}" value="{plugin.name}"/>
        {:else}
            <SelectItem disabled text="{plugin.name}" value="{plugin.name}"/>
        {/if}
    {/each}
</Select>

<script lang="ts">
    import {Select, SelectItem} from "carbon-components-svelte";
    import {pluginListStore} from "$lib/stores/plugins";
    import {onMount} from "svelte";

    export let profile;
    export let handleUpdate;

    let plugins = [];
    pluginListStore.subscribe((p) => plugins = p);
    onMount(async () => plugins = await pluginListStore.mount());

    /** handleSelectChange export selected profile name */
    const handleSelectChange = (e) => {
        profile.plugin = e.target.value;
        if (handleUpdate) {
            handleUpdate();
        }
    };
</script>
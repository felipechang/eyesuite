<StructuredListRow>
    <StructuredListCell>
        {profile.name}
    </StructuredListCell>
    <StructuredListCell>
        <PluginSelector handleUpdate="{handleUpdate}" bind:profile="{profile}"/>
    </StructuredListCell>
    <StructuredListCell>
        <Button on:click={()=>(open=!open)}>Edit</Button>
        <MappingFormModal
                handleUpdate="{handleModalUpdate}"
                bind:mapping="{mapping}"
                open="{open}"
                enabled
        />
    </StructuredListCell>
    <StructuredListCell>
        <TextArea on:change={handleUpdate} bind:value="{profile.template}"/>
    </StructuredListCell>
    <StructuredListCell>
        <TrashCan24 class="clickable" on:click={()=>handleDelete(profile)}/>
    </StructuredListCell>
</StructuredListRow>

<script lang="ts">
    import {Button, StructuredListCell, StructuredListRow, TextArea} from "carbon-components-svelte";
    import {TrashCan24} from "carbon-icons-svelte";
    import PluginSelector from "$lib/components/PluginSelector.svelte";
    import MappingFormModal from "$lib/components/MappingFormModal.svelte";
    import {onMount} from "svelte";

    let open = false;
    let mapping = [];

    export let handleUpdate;
    export let handleDelete;
    export let profile;

    onMount(() => {
        mapping = profile.mapping;
    });

    const handleModalUpdate = () => {
        profile.mapping = mapping;
        handleUpdate();
    }
</script>

<style>
    :global(svg.clickable) {
        cursor: pointer;
    }
</style>
<StructuredListRow>
    <StructuredListCell noWrap>
        <Checkbox checked="{plugin.enabled ? 'checked': ''}"
                  on:change={()=>handleEnableChange(plugin)}/>
    </StructuredListCell>
    <StructuredListCell>{plugin.name}</StructuredListCell>
    <StructuredListCell>{plugin.description}</StructuredListCell>
    <StructuredListCell>
        <Button on:click={()=>(open=!open)}>View</Button>
        <MappingFormModal
                enabled="{false}"
                handleUpdate="{handleUpdate}"
                mapping="{plugin.mapping}"
                open="{open}"
        />
    </StructuredListCell>
</StructuredListRow>


<script lang="ts">
    import {Button, Checkbox, StructuredListCell, StructuredListRow,} from "carbon-components-svelte";
    import MappingFormModal from "$lib/components/MappingFormModal.svelte";

    export let handleUpdate;
    export let plugin;

    let open = false;
    let mapping = [];

    /** handleEnableChange Persist change when plugin is enabled */
    const handleEnableChange = (p) => {
        plugin.enabled = !p.enabled;
        handleUpdate();
    }
</script>
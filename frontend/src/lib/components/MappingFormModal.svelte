<ComposedModal
        bind:open
        on:close={handleClose}
>
    <ModalBody hasForm>
        <Grid>
            <Row>
                <Column style="padding-bottom: 20px">
                    <div>
                        <span class="text">
                            Manage Mappings
                        </span>
                        {#if mapping.length !== 0 && enabled}
                            <span class="icon">
                            <AddFilled24 class="clickable" on:click={addNewLine}/>
                        </span>
                        {/if}
                    </div>
                    <p>{enabled ? "Add/Edit" : "View"} profile mappings</p>
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">Field Label</Column>
                <Column style="padding-bottom: 10px">Field Name</Column>
                <Column style="padding-bottom: 10px">Field Type</Column>
                {#if enabled}
                    <Column style="padding-bottom: 10px">
                        {mapping.length === 0 ? "Add" : "Remove"} Line
                    </Column>
                {/if}
            </Row>
            {#if mapping.length === 0 && enabled}
                <Row>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.name} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.id} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.type} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <AddAlt24 class="clickable" on:click={()=>handleAdd(map)}/>
                    </Column>
                </Row>
            {:else}
                {#each mapping as m}
                    <Row>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.name} required/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.id} required/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.type} required/>
                        </Column>
                        {#if enabled}
                            <Column style="padding-bottom: 10px">
                                <TrashCan24 class="clickable" on:click={()=>handleDelete(m)}/>
                            </Column>
                        {/if}
                    </Row>
                {/each}
            {/if}
        </Grid>
    </ModalBody>
</ComposedModal>

<script>
    import {Column, ComposedModal, Grid, ModalBody, Row, TextInput} from "carbon-components-svelte";
    import {AddAlt24, AddFilled24, TrashCan24} from "carbon-icons-svelte";

    export let open;
    export let enabled = false;
    export let handleUpdate;
    export let mapping;

    let map = {
        name: "",
        id: "",
        type: "",
    };

    const addNewLine = () => {
        mapping.push({
            id: "",
            name: "",
            type: "",
        });
        mapping = mapping;
        handleUpdate();
    }

    const handleClose = () => {
        mapping = mapping.filter((e) => {
            return e.id && e.name && e.type
        });
        handleUpdate();
    }

    const handleAdd = (m) => {
        if (m.name && m.id && m.type) {
            mapping.push(m);
            mapping = mapping;
        }
        handleUpdate();
    }

    const handleDelete = (m) => {
        mapping = mapping.filter((e) => {
            return e.name !== m.name;
        });
        handleUpdate();
    }
</script>

<style>
    :global(svg.clickable) {
        cursor: pointer;
    }

    .icon, .text {
        vertical-align: middle;
        display: inline-block;
    }

    .icon {
        padding-left: 10px;
    }

    .text {
        font-size: 2em;
        padding-bottom: 10px;
    }
</style>
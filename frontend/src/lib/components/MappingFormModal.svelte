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
                <Column style="padding-bottom: 10px">Label</Column>
                <Column style="padding-bottom: 10px">Name</Column>
                <Column style="padding-bottom: 10px">Value</Column>
                <Column style="padding-bottom: 10px">Placeholder</Column>
                <Column style="padding-bottom: 10px">Read-only</Column>
                <Column style="padding-bottom: 10px">Hidden</Column>
                {#if enabled}
                    <Column style="padding-bottom: 10px">
                        {mapping.length === 0 ? "Add" : "Remove"} Line
                    </Column>
                {/if}
            </Row>
            {#if mapping.length === 0 && enabled}
                <Row>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.labelText} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.name} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.value} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <TextInput bind:value={map.placeholder} required/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <Checkbox
                                checked="{map.readonly ? 'checked': ''}"
                                on:change={map.readonly = !map.readonly}/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <Checkbox
                                checked="{map.hidden ? 'checked': ''}"
                                on:change={map.readonly = !map.readonly}/>
                    </Column>
                    <Column style="padding-bottom: 10px">
                        <AddAlt24 class="clickable" on:click={()=>handleAdd(map)}/>
                    </Column>
                </Row>
            {:else}
                {#each mapping as m}
                    <Row>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.labelText} required/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.name} required/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.value} required/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <TextInput disabled="{!enabled}" bind:value={m.placeholder} required/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <Checkbox
                                    disabled="{!enabled}"
                                    checked="{m.readonly ? 'checked': ''}"
                                    on:change={m.readonly = !m.readonly}/>
                        </Column>
                        <Column style="padding-bottom: 10px">
                            <Checkbox
                                    disabled="{!enabled}"
                                    checked="{m.hidden ? 'checked': ''}"
                                    on:change={()=>m.hidden = !m.hidden}/>
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
    import {Checkbox, Column, ComposedModal, Grid, ModalBody, Row, TextInput} from "carbon-components-svelte";
    import {AddAlt24, AddFilled24, TrashCan24} from "carbon-icons-svelte";

    export let open;
    export let enabled = false;
    export let handleUpdate;
    export let mapping;

    let map = {
        labelText: "",
        name: "",
        value: "",
        placeholder: "",
        readonly: false,
        hidden: false,
    };

    const addNewLine = () => {
        mapping.push({
            labelText: "",
            name: "",
            value: "",
            placeholder: "",
            readonly: false,
            hidden: false,
        });
        mapping = mapping;
        handleUpdate();
    }

    const handleClose = () => {
        mapping = mapping.filter((e) => {
            return e.labelText && e.name
        });
        handleUpdate();
    }

    const handleAdd = (m) => {
        if (m.labelText && m.name) {
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
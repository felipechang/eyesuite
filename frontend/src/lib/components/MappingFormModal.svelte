<ComposedModal
        bind:open
        on:close={handleClose}
        size="lg"
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
            {#if mapping.length === 0 && enabled}
                <Row>
                    <Column style="min-width: 20%">
                        <TextInput
                                labelText="Label"
                                bind:value={map.labelText}
                                required
                        />
                    </Column>
                    <Column>
                        <TextInput
                                labelText="Name"
                                bind:value={map.name}
                                required
                        />
                    </Column>
                    <Column>
                        <TextInput
                                labelText="Value"
                                bind:value={map.value}
                                required
                        />
                    </Column>
                    <Column style="min-width: 25%">
                        <TextInput
                                labelText="Placeholder"
                                bind:value={map.placeholder}
                                required
                        />
                    </Column>
                    <Column style="max-width: 80px">
                        <Checkbox
                                labelText="Read Only"
                                checked="{map.readonly ? 'checked': ''}"
                                on:change={map.readonly = !map.readonly}/>
                    </Column>
                    <Column style="max-width: 80px">
                        <Checkbox
                                labelText="Hidde"
                                checked="{map.hidden ? 'checked': ''}"
                                on:change={map.hidden = !map.hidden}/>
                    </Column>
                    <Column>
                        <AddAlt24 class="clickable" on:click={()=>handleAdd(map)}/>
                    </Column>
                </Row>
            {:else}
                {#each mapping as m}
                    <Row>
                        <Column style="min-width: 20%">
                            <TextInput
                                    labelText="Label"
                                    disabled="{!enabled}"
                                    bind:value={m.labelText}
                                    required
                            />
                        </Column>
                        <Column>
                            <TextInput
                                    labelText="Name"
                                    disabled="{!enabled}"
                                    bind:value={m.name}
                                    required
                            />
                        </Column>
                        <Column>
                            <TextInput
                                    labelText="Value"
                                    disabled="{!enabled}"
                                    bind:value={m.value}
                                    required
                            />
                        </Column>
                        <Column style="min-width: 25%">
                            <TextInput
                                    labelText="Placeholder"
                                    disabled="{!enabled}"
                                    bind:value={m.placeholder}
                                    required
                            />
                        </Column>
                        <Column style="max-width: 80px">
                            <Checkbox
                                    labelText="Read Only"
                                    disabled="{!enabled}"
                                    checked="{m.readonly ? 'checked': ''}"
                                    on:change={m.readonly = !m.readonly}/>
                        </Column>
                        <Column style="max-width: 80px">
                            <Checkbox
                                    labelText="Hidden"
                                    disabled="{!enabled}"
                                    checked="{m.hidden ? 'checked': ''}"
                                    on:change={()=>m.hidden = !m.hidden}/>
                        </Column>
                        {#if enabled}
                            <Column>
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
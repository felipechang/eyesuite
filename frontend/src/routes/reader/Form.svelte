<Row>
    <Column style="padding-top: 40px">
        <FluidForm>
            <Grid>
                {#each fields as field}
                    <Row>
                        <Column style="padding-bottom: 10px">
                            <span style="display: {field.hidden ? 'none' : 'initial'}">
                                {#if field.list.length > 0}
                                    <Select style="height: 200px" labelText="{field.labelText}" size="xl"
                                            bind:selected="{field.value}">
                                        <SelectItem text="" value=""/>
                                        {#each field.list as list}
                                            <SelectItem text="{list}" value="{list}"/>
                                        {/each}
                                    </Select>
                                {:else}
                                    <TextInput
                                            labelText="{field.labelText}"
                                            name="{field.name}"
                                            placeholder="{field.placeholder}"
                                            bind:value="{field.value}"
                                            readonly="{field.readonly}"
                                    />
                                {/if}
                            </span>
                        </Column>
                    </Row>
                {/each}
                <Row>
                    {#if pluginLink !== ""}
                        <Column style="text-align: center; width: 300px; padding-top: 10px">
                            <Button
                                    style="width: 25%"
                                    icon={WatsonHealthMagnify16}
                                    on:click={handleImageRead}
                            >
                                Read
                            </Button>
                        </Column>
                    {/if}
                    <Column style="text-align: center; width: 300px; padding-top: 10px">
                        <Button
                                icon={Upload16}
                                on:click={handleImageSubmit}
                                style="width: 25%"
                        >
                            Submit
                        </Button>
                    </Column>
                </Row>
            </Grid>
        </FluidForm>
    </Column>
</Row>

<script lang="ts">
    import {Button, Column, FluidForm, Grid, Row, Select, SelectItem, TextInput} from "carbon-components-svelte";
    import {readerStore} from "$lib/stores/reader";
    import {Upload16, WatsonHealthMagnify16} from "carbon-icons-svelte";
    import {errorStore} from "$lib/stores/error";

    export let mappings = [];
    export let pluginLink;
    export let pluginField;
    export let template;
    export let files;
    export let success;
    export let loadMappings;
    export let updateSelectChange;

    let fields;
    $:fields = mappings.map((e) => {
        if (!e.value) {
            e.value = "";
        }
        return e;
    });

    const handleImageSubmit = async () => {
        if (!template || !files) {
            return;
        }
        fields.map((e) => template = template.replace(`{{.${e.name}}}`, e.value));
        const res = await readerStore.submit(template, files);
        if (!res.error) {
            success = res.msg;
            await loadMappings();
            updateSelectChange(localStorage.getItem("profile-selected"));
        } else {
            success = "";
            errorStore.set({
                title: "NetSuite Authorization Error:",
                error: res.msg
            });
        }
    }

    const handleImageRead = async () => {
        let res = await readerStore.read(pluginLink, files, fields);
        if (res.error) {
            errorStore.set({
                title: "Format error:",
                error: res.msg
            });
            return;
        }
        mappings = mappings.map((m) => {
            if (m.name === pluginField) {
                m.value = res.msg;
            }
            return m;
        });
    }
</script>

<style>
    :global(div.bx--select-input__wrapper) {
        height: 120px;
    }

</style>
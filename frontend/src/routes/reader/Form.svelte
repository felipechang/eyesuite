<Row>
    <Column style="padding-top: 40px">
        <FluidForm>
            <Grid>
                {#each fields as field}
                    <Row>
                        <Column style="padding-bottom: 10px">
                            <span style="display: {field.hidden ? 'none' : 'initial'}">
                               <TextInput
                                       labelText="{field.labelText}"
                                       name="{field.name}"
                                       placeholder="{field.placeholder}"
                                       bind:value="{field.value}"
                                       readonly="{field.readonly}"
                               />
                            </span>
                        </Column>
                    </Row>
                {/each}
                <Row>
                    <Column style="text-align: center; width: 300px; padding-top: 10px">
                        <Button
                                style="width: 25%"
                                icon={Upload16}
                                on:click={()=>readerStore.submit(fields)}
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
    import {Button, Column, FluidForm, Grid, Row, TextInput} from "carbon-components-svelte";
    import {readerStore} from "$lib/stores/reader";
    import {Upload16} from "carbon-icons-svelte";

    export let mappings = [];

    let fields;
    $:fields = mappings.map((e) => {
        if (!e.value) {
            e.value = "";
        }
        return e;
    });
</script>
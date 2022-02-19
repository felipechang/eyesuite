<Modal
        bind:open
        modalHeading="Add Profile"
        on:click:button--secondary={()=>(open=false)}
        on:close={handleClose}
        on:open
        on:submit={handleSubmit}
        primaryButtonText="Add"
        secondaryButtonText="Cancel"
>
    <FluidForm>
        <Grid>
            <Row>
                <Column style="padding-bottom: 10px">
                    <TextInput
                            labelText="Name"
                            bind:value={profile.name}
                            on:change={handleProfileChange}
                            required
                    />
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    Plugin:
                    <PluginSelector bind:profile/>
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    JSON Mapping:<br/>
                    <Button on:click={()=>(openModal=!openModal)}>Edit</Button>
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    XML Template:
                    <TextArea bind:value="{profile.template}"/>
                </Column>
            </Row>
        </Grid>
    </FluidForm>
</Modal>
<MappingFormModal
        handleUpdate="{handleModalUpdate}"
        bind:mapping="{mapping}"
        open="{openModal}"
        enabled
/>

<script lang="ts">
    import {Button, Column, FluidForm, Grid, Modal, Row, TextArea, TextInput} from "carbon-components-svelte";
    import PluginSelector from "$lib/components/PluginSelector.svelte";
    import {profileListStore} from "$lib/stores/profiles";
    import {errorStore} from "$lib/stores/error";
    import MappingFormModal from "$lib/components/MappingFormModal.svelte";

    export let open;
    export let profiles;

    let openModal = false;
    let mapping = [];

    let profile = {
        name: "",
        plugin: "",
        mapping: [],
        template: "<>",
    };

    /** handleClose Reset values and close modal */
    const handleClose = () => {
        profile.name = "";
        profile.plugin = "";
        profile.mapping = [];
        profile.template = "<>";
    }

    /** handleSubmit Add profile and close modal */
    const handleSubmit = () => {
        profiles.push({
            name: profile.name,
            plugin: profile.plugin,
            mapping: profile.mapping,
            template: profile.template,
        });
        profileListStore.persist(profiles);
        open = false;
    }

    /** handleProfileChange Show a duplicate error if name is repeated */
    const handleProfileChange = () => {
        profile.name = profile.name.trim();
        for (let i = 0; i < profiles.length; i++) {
            if (profiles[i].name.trim() === profile.name) {
                errorStore.set({
                    title: "Duplicate record",
                    error: "Profiles names must be unique",
                });
                profile.name = "";
                return;
            }
        }
    }

    const handleModalUpdate = () => {
        profile.mapping = mapping;
        profileListStore.persist(profiles);
    }
</script>
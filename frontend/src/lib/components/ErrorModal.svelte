<ComposedModal danger on:click:button--primary={handleClose} open="{open}" preventCloseOnClickOutside>
    <ModalHeader label="Error" title="{title}"/>
    <ModalBody hasForm>{error}</ModalBody>
    <ModalFooter primaryButtonDisabled={false} primaryButtonText="Ok"/>
</ComposedModal>

<script lang="ts">
    import {ComposedModal, ModalBody, ModalFooter, ModalHeader} from "carbon-components-svelte";
    import {errorStore} from "$lib/stores/error";

    let open = false;
    let title = "";
    let error = "";

    errorStore.subscribe((e) => {
        if (e.title && e.error) {
            title = e.title;
            error = e.error;
            open = true;
        }
    });

    /** handleClose Reset values and close modal */
    const handleClose = () => {
        title = "";
        error = "";
        open = false;
    }

</script>
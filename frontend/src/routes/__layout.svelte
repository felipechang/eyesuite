<Header/>
<Content>
    <slot/>
</Content>
<ErrorModal/>

<script lang="ts">
    import {Content} from "carbon-components-svelte";
    import "carbon-components-svelte/css/g10.css";
    import Header from "$lib/components/Header.svelte";
    import ErrorModal from "$lib/components/ErrorModal.svelte";
</script>

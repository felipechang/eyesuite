{#if logged}
    <Header
            handleLogOut="{()=>logged = false}"
            isAdministrator="{isAdministrator}"
    />
    <Content>
        <slot isAdministrator="{isAdministrator}"/>
    </Content>
{:else }
    <Login
            handleLogIn="{()=>logged = true}"
            bind:isAdministrator="{isAdministrator}"
    />
{/if}
<ErrorModal/>

<script lang="ts">
    import {Content} from "carbon-components-svelte";
    import "carbon-components-svelte/css/g10.css";
    import Header from "$lib/components/Header.svelte";
    import ErrorModal from "$lib/components/ErrorModal.svelte"
    import Login from "$lib/components/Login.svelte";
    import {onMount} from "svelte";

    let logged = false;
    let isAdministrator = false;

    onMount(() => logged = !!localStorage.getItem("auth"));
</script>

<Header
        bind:isSideNavOpen
        company="EyeSuite"
        persistentHamburgerMenu
        platformName="NetSuite Reader"
>
    <svelte:fragment slot="skip-to-content">
        <SkipToContent/>
    </svelte:fragment>
    <HeaderUtilities>
        <HeaderGlobalAction aria-label="Settings" icon={Logout24} on:click={handleServerLogout}/>
    </HeaderUtilities>
</Header>

<SideNav bind:isOpen={isSideNavOpen}>
    <SideNavItems>
        <SideNavLink href="/" on:click={handleClick} text="Reader"/>
        <SideNavLink href="/help" on:click={handleClick} text="Help"/>
        {#if isAdministrator}
            <SideNavMenu text="Settings">
                <SideNavMenuItem href="/settings/connection" on:click={handleClick} text="Connection"/>
                <SideNavMenuItem href="/settings/plugins" on:click={handleClick} text="Plugins"/>
                <SideNavMenuItem href="/settings/profiles" on:click={handleClick} text="Profiles"/>
                <SideNavMenuItem href="/settings/users" on:click={handleClick} text="Users"/>
            </SideNavMenu>
        {/if}
    </SideNavItems>
</SideNav>

<script lang="ts">
    import {
        Header,
        HeaderGlobalAction,
        HeaderUtilities,
        SideNav,
        SideNavItems,
        SideNavLink,
        SideNavMenu,
        SideNavMenuItem,
        SkipToContent,
    } from "carbon-components-svelte";
    import "carbon-components-svelte/css/g10.css";
    import {Logout24} from "carbon-icons-svelte";
    import {loginStore} from "$lib/stores/login";
    import {errorStore} from "$lib/stores/error";

    export let handleLogOut;
    export let isAdministrator;

    let isSideNavOpen = false;

    /** handleClick Close modal */
    const handleClick = () => {
        isSideNavOpen = false;
    }

    const handleServerLogout = async () => {
        const res = await loginStore.logout();
        if (res.error) {
            errorStore.set({title: "Logout Error", error: res.data.toUpperCase()});
        }
        localStorage.removeItem("auth");
        localStorage.removeItem("user");
        localStorage.removeItem("profile-selected");
        handleLogOut();
    }
</script>

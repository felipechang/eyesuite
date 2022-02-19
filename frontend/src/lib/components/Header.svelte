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
        <HeaderGlobalAction aria-label="Settings" icon={Logout24} on:click={handleLogOut}/>
    </HeaderUtilities>
</Header>

<SideNav bind:isOpen={isSideNavOpen}>
    <SideNavItems>
        <SideNavLink href="/" on:click={handleClick} text="Reader"/>
        <SideNavLink href="/help" on:click={handleClick} text="Help"/>
        <SideNavMenu text="Settings">
            <SideNavMenuItem href="/settings/connection" on:click={handleClick} text="Connection"/>
            <SideNavMenuItem href="/settings/plugins" on:click={handleClick} text="Plugins"/>
            <SideNavMenuItem href="/settings/profiles" on:click={handleClick} text="Profiles"/>
            <SideNavMenuItem href="/settings/users" on:click={handleClick} text="Users"/>
        </SideNavMenu>
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

    let isSideNavOpen = false;

    /** handleClick Close modal */
    const handleClick = () => {
        isSideNavOpen = false;
    }

    /** handleLogOut Close session */
    const handleLogOut = () => {
        console.log("log out");
    }
</script>

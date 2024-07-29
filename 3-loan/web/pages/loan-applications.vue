<template>
    <div class="flex h-screen overflow">
        <SideMenu />
        <div class="flex-1 flex flex-col">
            <HeaderMenu />
            <main class="p-6">
                <header class="mb-6">
                    <h2 class="text-2xl font-bold">Loan Applications</h2>
                </header>
                <DataTable />
            </main>
        </div>
    </div>
</template>
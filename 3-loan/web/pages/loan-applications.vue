<template>
    <div class="flex h-screen overflow">
        <SideMenu />
        <div class="flex-1 flex flex-col">
            <HeaderMenu />
            <main class="p-6">
                <header class="mb-6">
                    <h2 class="text-2xl font-bold">Loan Applications</h2>
                </header>
                <div class="rounded-sm border border-stroke bg-white shadow-default">
                    <div class="px-8 py-5 flex border-b border-gray-100">
                        <div class="w-96">
                            <input placeholder="Search..." class="w-full py-4 px-6 outline-none border-stroke border rounded-lg bg-gray-50" />
                        </div>
                        <div></div>
                    </div>
                    <table class="table">
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>Requested At</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Ivan Ivanov Ivanovich</td>
                                <td>January 1, 2024 9:00 AM</td>
                            </tr>
                        </tbody>
                    </table>
                    <div class="px-8 py-5 flex border-b border-gray-100">
                        <div>Showing 1 to 10 of 26 entries</div>
                        <div></div>
                    </div>
                </div>
            </main>
        </div>
    </div>
</template>
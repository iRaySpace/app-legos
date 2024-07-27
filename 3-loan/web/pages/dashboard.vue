<template>
    <div class="flex h-screen overflow">
        <SideMenu />
        <div class="flex-1 flex flex-col">
            <header class="w-full border-b border-gray-200 bg-white">
                <div class="px-6 py-4 flex justify-between">
                    <div></div>
                    <div class="dropdown dropdown-end">
                        <a href="#" role="button" class="flex items-center gap-4">
                            <span>
                                <span class="block text-sm font-medium text-black">Ivan Ivanov Ivanovich</span>
                                <span class="block text-xs font-medium">Software Engineer</span>
                            </span>
                            <span class="h-12 w-12 rounded-full">
                                <img src="/user-01.png" alt="User">
                            </span>
                        </a>
                        <ul class="menu dropdown-content bg-base-100 rounded-box z-[1] mt-4 w-52 p-2 shadow">
                            <li><a>My Profile</a></li>
                            <li><a>Logout</a></li>
                        </ul>
                    </div>
                </div>
            </header>
        </div>
    </div>
</template>
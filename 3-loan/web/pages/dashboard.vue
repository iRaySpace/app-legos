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
                        <div class="dropdown-content bg-white z-[1] mt-4 w-60 rounded-sm shadow">
                            <ul class="px-6 py-6 flex flex-col gap-5">
                                <li><a href="#" class="flex gap-3.5 hover:text-primary text-sm font-medium text-gray-600">
                                        <IconsPerson /> My Profile
                                    </a></li>
                                <li><a href="#" class="flex gap-3.5 hover:text-primary text-sm font-medium text-gray-600">
                                        <IconsExit /> Log Out
                                    </a></li>
                            </ul>
                        </div>
                    </div>
                </div>
            </header>
        </div>
</div></template>
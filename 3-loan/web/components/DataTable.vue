<template>
    <div class="rounded-sm border border-stroke bg-white shadow-default">
        <div class="px-8 py-5 flex border-b border-gray-100 justify-between items-center">
            <div class="w-96 relative">
                <input placeholder="Search..."
                    class="w-full py-3 pl-6 pr-10 outline-none border-stroke border rounded-lg bg-gray-50" />
                <span class="absolute right-4 top-4">
                    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true"
                        role="img" class="text-base-content/60" font-size="15" width="1em" height="1em" viewBox="0 0 24 24">
                        <g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"
                            stroke-width="2">
                            <circle cx="11" cy="11" r="8"></circle>
                            <path d="m21 21l-4.3-4.3"></path>
                        </g>
                    </svg>
                </span>
            </div>
            <div>
                <a href="#" class="btn btn-primary text-white">
                    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true"
                        role="img" font-size="16" width="1em" height="1em" viewBox="0 0 24 24">
                        <path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"
                            stroke-width="2" d="M5 12h14m-7-7v14"></path>
                    </svg>
                    Add New
                </a>
            </div>
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
        <div class="px-8 py-5 flex border-b border-gray-100 justify-between items-center">
            <div>Showing 1 to 10 of 26 entries</div>
            <div>
                <div class="join">
                    <button class="join-item btn">«</button>
                    <button class="join-item btn">Page 1</button>
                    <button class="join-item btn">»</button>
                </div>
            </div>
        </div>
    </div>
</template>
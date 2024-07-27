<template>
    <main class="h-screen py-12">
        <header class="flex flex-col justify-center">
            <img alt="Loan Company" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&amp;shade=600" class="h-10">
            <h2 class="font-bold text-3xl text-center mt-6">
                Sign in to your account
            </h2>
        </header>
        <div class="max-w-[480px] mx-auto mt-9 p-12 bg-white border border-gray-200 rounded-lg">
            <form>
                <div class="mb-4">
                    <label class="font-medium block mb-2.5">Email address</label>
                    <div class="relative">
                        <input name="email" type="email" placeholder="Enter your email"
                            class="pr-10 pl-6 py-4 outline-none border-stroke border rounded-lg w-full" required>
                        <span class="absolute right-4 top-4">
                            <IconsMail />
                        </span>
                    </div>
                </div>
                <div class="mb-6">
                    <label class="font-medium block mb-2.5">Password</label>
                    <div class="relative">
                        <input name="password" type="password" placeholder="Enter your password"
                            class="pr-10 pl-6 py-4 outline-none border-stroke border rounded-lg w-full" required>
                        <span class="absolute right-4 top-4">
                            <IconsLock />
                        </span>
                    </div>
                </div>
                <div class="mb-5">
                    <input type="submit" class="w-full cursor-pointer rounded-lg border border-primary bg-primary p-4 font-medium text-white hover:bg-opacity-90">
                </div>
            </form>
        </div>
    </main>
</template>
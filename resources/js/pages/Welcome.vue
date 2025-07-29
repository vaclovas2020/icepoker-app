<script setup lang="ts">
import { Head, Link } from '@inertiajs/vue3';
</script>

<template>

    <Head title="ICE Poker for Small IT Teams" />
    <div
        class="flex min-h-screen flex-col items-center bg-[#FDFDFC] p-6 text-[#1b1b18] lg:justify-center lg:p-8 dark:bg-[#0a0a0a]">
        <header class="mb-6 w-full max-w-[335px] text-sm not-has-[nav]:hidden lg:max-w-4xl">
            <nav class="flex items-center justify-end gap-4">
                <Link v-if="$page.props.auth.user" :href="route('dashboard')"
                    class="inline-block rounded-sm border border-[#19140035] px-5 py-1.5 text-sm leading-normal text-[#1b1b18] hover:border-[#1915014a] dark:border-[#3E3E3A] dark:text-[#EDEDEC] dark:hover:border-[#62605b]">
                Dashboard
                </Link>
                <template v-else>
                    <Link :href="route('login')"
                        class="inline-block rounded-sm border border-transparent px-5 py-1.5 text-sm leading-normal text-[#1b1b18] hover:border-[#19140035] dark:text-[#EDEDEC] dark:hover:border-[#3E3E3A]">
                    Log in
                    </Link>
                    <Link :href="route('register')"
                        class="inline-block rounded-sm border border-[#19140035] px-5 py-1.5 text-sm leading-normal text-[#1b1b18] hover:border-[#1915014a] dark:border-[#3E3E3A] dark:text-[#EDEDEC] dark:hover:border-[#62605b]">
                    Register
                    </Link>
                </template>
            </nav>
        </header>
        <div
            class="flex w-full items-center justify-center opacity-100 transition-opacity duration-750 lg:grow starting:opacity-0">
            <main
                class="flex w-full max-w-[335px] flex-col-reverse overflow-hidden rounded-lg lg:max-w-4xl lg:flex-row">
                <div
                    class="flex-1 rounded-br-lg rounded-bl-lg bg-white p-6 pb-12 text-[13px] leading-[20px] shadow-[inset_0px_0px_0px_1px_rgba(26,26,0,0.16)] lg:rounded-tl-lg lg:rounded-br-none lg:p-20 dark:bg-[#161615] dark:text-[#EDEDEC] dark:shadow-[inset_0px_0px_0px_1px_#fffaed2d]">
                    <h1 class="mb-1 font-medium">ICE Poker for Small IT Teams</h1>
                    <p class="mb-2 text-[#706f6c] dark:text-[#A1A09A]">
                        Think of ICE scoring like a poker game—you and your team assign values (Impact, Confidence,
                        Ease) to
                        each task and compare scores to decide what to tackle next. It’s a lightweight, fast
                        decision-making
                        framework that works well for lean teams with limited resources.
                    </p>
                    <ul class="mb-4 flex flex-col lg:mb-6">
                        <li
                            class="relative flex items-center gap-4 py-2 before:absolute before:top-1/2 before:bottom-0 before:left-[0.4rem] before:border-l before:border-[#e3e3e0] dark:before:border-[#3E3E3A]">
                            <span class="relative bg-white py-1 dark:bg-[#161615]">
                                <span
                                    class="flex h-3.5 w-3.5 items-center justify-center rounded-full border border-[#e3e3e0] bg-[#FDFDFC] shadow-[0px_0px_1px_0px_rgba(0,0,0,0.03),0px_1px_2px_0px_rgba(0,0,0,0.06)] dark:border-[#3E3E3A] dark:bg-[#161615]">
                                    <span class="h-1.5 w-1.5 rounded-full bg-[#dbdbd7] dark:bg-[#3E3E3A]" />
                                </span>
                            </span>
                            <span>
                                <strong>I - Impact:</strong> How much the task will positively affect the goal if it
                                succeeds
                                (scale:
                                1-10).
                            </span>
                        </li>
                        <li
                            class="relative flex items-center gap-4 py-2 before:absolute before:top-1/2 before:bottom-0 before:left-[0.4rem] before:border-l before:border-[#e3e3e0] dark:before:border-[#3E3E3A]">
                            <span class="relative bg-white py-1 dark:bg-[#161615]">
                                <span
                                    class="flex h-3.5 w-3.5 items-center justify-center rounded-full border border-[#e3e3e0] bg-[#FDFDFC] shadow-[0px_0px_1px_0px_rgba(0,0,0,0.03),0px_1px_2px_0px_rgba(0,0,0,0.06)] dark:border-[#3E3E3A] dark:bg-[#161615]">
                                    <span class="h-1.5 w-1.5 rounded-full bg-[#dbdbd7] dark:bg-[#3E3E3A]" />
                                </span>
                            </span>
                            <span>
                                <strong>C - Confidence:</strong> How sure you are that this task will succeed
                                and
                                deliver the expected impact (scale: 1-10).
                            </span>
                        </li>
                        <li
                            class="relative flex items-center gap-4 py-2 before:absolute before:top-1/2 before:bottom-0 before:left-[0.4rem] before:border-l before:border-[#e3e3e0] dark:before:border-[#3E3E3A]">
                            <span class="relative bg-white py-1 dark:bg-[#161615]">
                                <span
                                    class="flex h-3.5 w-3.5 items-center justify-center rounded-full border border-[#e3e3e0] bg-[#FDFDFC] shadow-[0px_0px_1px_0px_rgba(0,0,0,0.03),0px_1px_2px_0px_rgba(0,0,0,0.06)] dark:border-[#3E3E3A] dark:bg-[#161615]">
                                    <span class="h-1.5 w-1.5 rounded-full bg-[#dbdbd7] dark:bg-[#3E3E3A]" />
                                </span>
                            </span>
                            <span>
                                <strong>E - Ease:</strong> How easy it is to complete the task in terms
                                of time,
                                effort, and complexity (scale: 1-10).
                            </span>
                        </li>
                    </ul>
                    <ul class="flex gap-3 text-sm leading-normal">
                        <li>
                            <a :href="route('register')"
                                class="inline-block rounded-sm border border-black bg-[#1b1b18] px-5 py-1.5 text-sm leading-normal text-white hover:border-black hover:bg-black dark:border-[#eeeeec] dark:bg-[#eeeeec] dark:text-[#1C1C1A] dark:hover:border-white dark:hover:bg-white">
                                Get started
                            </a>
                        </li>
                    </ul>
                </div>
                <div
                    class="relative -mb-px aspect-335/376 w-full shrink-0 overflow-hidden rounded-t-lg bg-[#1dffc6] lg:mb-0 lg:-ml-px lg:aspect-auto lg:w-[438px] lg:rounded-t-none lg:rounded-r-lg">

                    <svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="512" height="512">
                        <path
                            d="M0 0 C1.13501953 0.35505615 1.13501953 0.35505615 2.29296875 0.71728516 C21.59537012 6.92338407 40.12202175 16.35715982 56 29 C56.54269531 29.43135254 57.08539063 29.86270508 57.64453125 30.30712891 C76.6183939 45.52065636 91.9678133 63.22594731 103 85 C103.32194336 85.63147949 103.64388672 86.26295898 103.97558594 86.91357422 C121.81425248 122.344416 125.70564778 165.16152971 113.3671875 203.03125 C104.69608835 228.51265406 91.64667786 249.6742322 73 269 C72.26394531 269.78761719 71.52789062 270.57523437 70.76953125 271.38671875 C59.80007648 282.83571252 46.67133991 291.17043762 33 299 C32.36288086 299.36867188 31.72576172 299.73734375 31.06933594 300.1171875 C-8.92785417 322.86278358 -61.08013184 327.19410962 -105.07373047 315.11181641 C-130.31914941 307.7697363 -154.15722043 296.47027332 -174 279 C-174.85078125 278.25878906 -175.7015625 277.51757812 -176.578125 276.75390625 C-179.08471581 274.53565027 -181.54854412 272.27893116 -184 270 C-184.71542969 269.3503125 -185.43085938 268.700625 -186.16796875 268.03125 C-198.22005421 256.61579565 -207.5054399 242.7237268 -215 228 C-215.31662598 227.38979004 -215.63325195 226.77958008 -215.95947266 226.15087891 C-234.16411932 190.73975158 -237.2578657 148.09452734 -225.46972656 110.17773438 C-222.85433386 102.19523882 -219.66587249 94.55243949 -216 87 C-215.70399902 86.37818848 -215.40799805 85.75637695 -215.10302734 85.11572266 C-207.60192609 69.53401853 -196.93455615 56.40221117 -185 44 C-184.38382813 43.32066406 -183.76765625 42.64132813 -183.1328125 41.94140625 C-138.30318535 -5.97648411 -60.74154058 -19.10461147 0 0 Z M-83 30 C-81.69983484 32.89631874 -80.36489577 34.85194306 -78.0234375 36.98828125 C-77.43433594 37.53033203 -76.84523437 38.07238281 -76.23828125 38.63085938 C-75.31595703 39.46327148 -75.31595703 39.46327148 -74.375 40.3125 C-73.78332031 40.85583984 -73.19164062 41.39917969 -72.58203125 41.95898438 C-69.51244287 44.76669642 -66.4017808 47.51620355 -63.24609375 50.2265625 C-57.84164093 54.94401969 -52.76093296 59.9951291 -47.6875 65.0625 C-47.17765564 65.57166962 -46.66781128 66.08083923 -46.14251709 66.60543823 C-38.46366061 74.30043881 -31.21928988 82.06045634 -25 91 C-24.60184082 91.56444824 -24.20368164 92.12889648 -23.79345703 92.71044922 C-15.73287569 104.23518808 -10.29947665 116.26640253 -7 130 C1.74872387 130.98223271 10.45122416 131.23893122 19.25 131.25 C20.31182999 131.25250763 20.31182999 131.25250763 21.39511108 131.25506592 C41.06959652 131.25781295 60.62709102 129.46572422 80 126 C73.47224832 95.20338229 51.09403195 68.42054425 25.26049805 51.42285156 C-6.95920515 31.00205363 -45.37532423 23.24685307 -83 30 Z M-161.38867188 71.32714844 C-175.89190364 86.38413996 -193.25487719 111.35556112 -193 133 C-186.72381854 131.35081723 -180.5332099 129.56019614 -174.37963867 127.50073242 C-135.31502153 114.49128282 -97.71030715 113.49138675 -57.390625 121.44921875 C-54.18837063 121.96940062 -51.23652982 122.10424406 -48 122 C-54.40896988 103.25548594 -77.27178451 84.1871378 -92.0234375 71.6953125 C-94.50860194 69.5637683 -96.90440146 67.34316401 -99.3125 65.125 C-103.908689 60.93041975 -108.63622842 56.95007628 -113.484375 53.05078125 C-115.71330919 51.23371533 -117.85995941 49.3572423 -120 47.4375 C-130.26512071 39.09708942 -153.4732994 63.44511388 -161.38867188 71.32714844 Z M-195 173 C-191.48320656 196.19849707 -182.29719743 217.10365301 -167 235 C-166.56429687 235.51401367 -166.12859375 236.02802734 -165.6796875 236.55761719 C-141.23526505 265.09397878 -105.31345441 282.0211021 -68 285 C-61.99510544 285.26506985 -56.00567348 285.19583718 -50 285 C-50.556875 283.86046875 -51.11375 282.7209375 -51.6875 281.546875 C-64.95216607 254.157227 -77.60206488 225.57337496 -67.68457031 194.96777344 C-64.1422999 184.78564095 -59.34243813 175.34557126 -53.75708008 166.13867188 C-52.76525565 164.46687983 -51.87351915 162.73653927 -51 161 C-51.33 160.01 -51.66 159.02 -52 158 C-56.92463521 157.15842167 -61.85018457 156.32248196 -66.77661133 155.49145508 C-68.44776384 155.20877336 -70.11866965 154.9246287 -71.78930664 154.63891602 C-113.70112386 147.47589667 -156.96377058 152.52393783 -195 173 Z M73.09765625 162.62109375 C71.38932617 162.81735352 71.38932617 162.81735352 69.64648438 163.01757812 C67.2922931 163.2934098 64.93812566 163.56944511 62.58398438 163.84570312 C49.32404761 165.36855301 36.20769717 166.16816561 22.86914062 166.13037109 C20.6340451 166.12502165 18.39915192 166.13034958 16.1640625 166.13671875 C6.40752455 166.14073213 -3.27107356 165.7240208 -13 165 C-13.25579834 165.49177734 -13.51159668 165.98355469 -13.77514648 166.49023438 C-14.9932585 168.82822795 -16.21491927 171.16434002 -17.4375 173.5 C-17.88689941 174.35988525 -18.33629883 175.21977051 -18.79931641 176.10571289 C-21.34456203 180.94757176 -23.96635516 185.72472161 -26.72119141 190.45019531 C-34.61145769 204.04028578 -38.23967038 215.56398975 -34.4375 231.25 C-30.42198627 245.57181637 -23.27573796 258.68886394 -16.35546875 271.7734375 C-15.94965576 272.54091309 -15.54384277 273.30838867 -15.12573242 274.09912109 C-14.0201477 276.14525412 -14.0201477 276.14525412 -12 278 C15.43669458 274.02603432 42.71471793 251.82643152 59.20922852 230.39111328 C73.38760412 211.13547288 84.54115617 186.21673876 84 162 C80.30936873 161.93277462 76.76249279 162.1906375 73.09765625 162.62109375 Z "
                            fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"
                            transform="translate(311,103)" />
                    </svg>
                </div>
            </main>
        </div>
        <div class="hidden h-14.5 lg:block"></div>
    </div>
</template>

import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const skylarksTheme: CustomThemeConfig = {
    name: 'skylarksTheme',
    properties: {
        // =~= Theme Properties =~=
        "--theme-font-family-base": `system-ui`,
        "--theme-font-family-heading": `system-ui`,
        "--theme-font-color-base": "var(--color-secondary-900)",
        "--theme-font-color-dark": "var(--color-tertiary-50)",
        "--theme-rounded-base": "12px",
        "--theme-rounded-container": "12px",
        "--theme-border-base": "1px",
        // =~= Theme On-X Colors =~=
        "--on-primary": "255 255 255",
        "--on-secondary": "255 255 255",
        "--on-tertiary": "var(--color-secondary-800)",
        "--on-success": "0 0 0",
        "--on-warning": "0 0 0",
        "--on-error": "255 255 255",
        "--on-surface": "255 255 255",
        // =~= Theme Colors  =~=
        // primary | #BA0C2F
        "--color-primary-50": "245 219 224", // #f5dbe0
        "--color-primary-100": "241 206 213", // #f1ced5
        "--color-primary-200": "238 194 203", // #eec2cb
        "--color-primary-300": "227 158 172", // #e39eac
        "--color-primary-400": "207 85 109", // #cf556d
        "--color-primary-500": "186 12 47", // #BA0C2F
        "--color-primary-600": "167 11 42", // #a70b2a
        "--color-primary-700": "140 9 35", // #8c0923
        "--color-primary-800": "112 7 28", // #70071c
        "--color-primary-900": "91 6 23", // #5b0617
        // secondary | #041E42
        "--color-secondary-50": "217 221 227", // #d9dde3
        "--color-secondary-100": "205 210 217", // #cdd2d9
        "--color-secondary-200": "192 199 208", // #c0c7d0
        "--color-secondary-300": "155 165 179", // #9ba5b3
        "--color-secondary-400": "79 98 123", // #4f627b
        "--color-secondary-500": "4 30 66", // #041E42
        "--color-secondary-600": "4 27 59", // #041b3b
        "--color-secondary-700": "3 23 50", // #031732
        "--color-secondary-800": "2 18 40", // #021228
        "--color-secondary-900": "2 15 32", // #020f20
        // tertiary | #CEB888
        "--color-tertiary-50": "248 244 237", // #f8f4ed
        "--color-tertiary-100": "245 241 231", // #f5f1e7
        "--color-tertiary-200": "243 237 225", // #f3ede1
        "--color-tertiary-300": "235 227 207", // #ebe3cf
        "--color-tertiary-400": "221 205 172", // #ddcdac
        "--color-tertiary-500": "206 184 136", // #CEB888
        "--color-tertiary-600": "185 166 122", // #b9a67a
        "--color-tertiary-700": "155 138 102", // #9b8a66
        "--color-tertiary-800": "124 110 82", // #7c6e52
        "--color-tertiary-900": "101 90 67", // #655a43
        // success | #84cc16
        "--color-success-50": "237 247 220", // #edf7dc
        "--color-success-100": "230 245 208", // #e6f5d0
        "--color-success-200": "224 242 197", // #e0f2c5
        "--color-success-300": "206 235 162", // #ceeba2
        "--color-success-400": "169 219 92", // #a9db5c
        "--color-success-500": "132 204 22", // #84cc16
        "--color-success-600": "119 184 20", // #77b814
        "--color-success-700": "99 153 17", // #639911
        "--color-success-800": "79 122 13", // #4f7a0d
        "--color-success-900": "65 100 11", // #41640b
        // warning | #EAB308
        "--color-warning-50": "252 244 218", // #fcf4da
        "--color-warning-100": "251 240 206", // #fbf0ce
        "--color-warning-200": "250 236 193", // #faecc1
        "--color-warning-300": "247 225 156", // #f7e19c
        "--color-warning-400": "240 202 82", // #f0ca52
        "--color-warning-500": "234 179 8", // #EAB308
        "--color-warning-600": "211 161 7", // #d3a107
        "--color-warning-700": "176 134 6", // #b08606
        "--color-warning-800": "140 107 5", // #8c6b05
        "--color-warning-900": "115 88 4", // #735804
        // error | #D41976
        "--color-error-50": "249 221 234", // #f9ddea
        "--color-error-100": "246 209 228", // #f6d1e4
        "--color-error-200": "244 198 221", // #f4c6dd
        "--color-error-300": "238 163 200", // #eea3c8
        "--color-error-400": "225 94 159", // #e15e9f
        "--color-error-500": "212 25 118", // #D41976
        "--color-error-600": "191 23 106", // #bf176a
        "--color-error-700": "159 19 89", // #9f1359
        "--color-error-800": "127 15 71", // #7f0f47
        "--color-error-900": "104 12 58", // #680c3a
        // surface | #242c46
        "--color-surface-50": "222 223 227", // #dedfe3
        "--color-surface-100": "211 213 218", // #d3d5da
        "--color-surface-200": "200 202 209", // #c8cad1
        "--color-surface-300": "167 171 181", // #a7abb5
        "--color-surface-400": "102 107 126", // #666b7e
        "--color-surface-500": "36 44 70", // #242c46
        "--color-surface-600": "32 40 63", // #20283f
        "--color-surface-700": "27 33 53", // #1b2135
        "--color-surface-800": "22 26 42", // #161a2a
        "--color-surface-900": "18 22 34", // #121622
    }
}
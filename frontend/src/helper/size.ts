export const kb = 1024;
export const mb = kb * 1024;
export const gb = mb * 1024;
export const tb = gb * 1024;
export const tenGb = gb * 10;

export function toShortSize(size: number): string {
    if (size < kb) {
        return `${size} B`
    }
    if (size < mb) {
        const s = size / kb < 100 ? (size / kb).toFixed(1) : (size / kb).toFixed(0)
        return `${s} KB`
    }
    if (size < gb) {
        const s = size / mb < 100 ? (size / mb).toFixed(1) : (size / mb).toFixed(0)
        return `${s} MB`
    }
    if (size < tb) {
        const s = size / gb < 100 ? (size / gb).toFixed(1) : (size / gb).toFixed(0)
        return `${s} GB`
    }
    const s = size / tb < 100 ? (size / tb).toFixed(1) : (size / tb).toFixed(0)
    return `${s} TB`
}
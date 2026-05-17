type Params = {
    route: string;
    method: string;
    error: {
        message: string;
        alert: string;
    }
}

export async function apiFetch<T>(params: Params): Promise<T> {
    try {
        const response = await fetch(params.route, {
            method: params.method,
            headers: {
                'Content-Type': 'application/json',
            },
        })

        if (!response.ok) {
            throw new Error('サーバーエラーが発生しました')
        }

        const res: T = await response.json()

        return res

    } catch (error: any) {
        alert(params.error.alert)
        throw new Error(params.error.message)
    }
}
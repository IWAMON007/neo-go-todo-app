type Params = {
    route: string;
    method: string;
    body?: object;
    error: {
        message: string;
    };
};

export async function apiFetch<T>(params: Params): Promise<T> {
    try {
        const response = await fetch(params.route, {
            method: params.method,
            headers: {
                'Content-Type': 'application/json',
            },
            body: params.body ? JSON.stringify(params.body) : undefined,
        });

        if (!response.ok) {
            throw new Error(params.error.message);
        }

        const res: T = await response.json();

        return res;
    } catch (error: unknown) {
        if (error instanceof Error) {
            throw error;
        }

        throw new Error('予期せぬエラーが発生しました。');
    }
}

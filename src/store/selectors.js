export const getMe = (state) => {
    return state.appStore.me;
}

export const getCredits = (state) => {
    return state.appStore.credits;
}

export const getMeFetched = (state) => {
    return state.appStore.meFetched;
}

export const getCreditsFetched = (state) => {
    return state.appStore.creditsFetched;
}
export const SET_ME = 'SET-ME';
export const SET_CREDITS = 'SET-CREDITS';
export const SET_ME_FETCHED = 'SET-ME-FETCHED';
export const SET_CREDITS_FETCHED = 'SET-CREDITS-FETCHED';

export const onSetMeActionCreator = (data) => {
    return{
        type: SET_ME,
        data: data
    }
}

export const onSetCreditsActionCreator = (data) => {
    return{
        type: SET_CREDITS,
        data: data
    }
}

export const onSetMeFetchedActionCreator = (data) => {
    return{
        type: SET_ME_FETCHED,
        data: data
    }
}

export const onSetCreditsFetchedActionCreator = (data) => {
    return{
        type: SET_CREDITS_FETCHED,
        data: data
    }
}
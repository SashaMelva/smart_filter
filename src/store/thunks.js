import {
    onSetCreditsActionCreator,
    onSetCreditsFetchedActionCreator,
    onSetMeActionCreator,
    onSetMeFetchedActionCreator
} from "./action_types";
import {userAPI} from "./api";

export const getMeThunkCreator = () => {
    return async (dispatch) => {
        const data = await userAPI.getMe();
        dispatch(onSetMeActionCreator(data));
        dispatch(onSetMeFetchedActionCreator(true));
    }
}

export const getCreditsThunkCreator = () => {
    return async (dispatch) => {
        const data = await userAPI.getCredits();
        dispatch(onSetCreditsActionCreator(data));
        dispatch(onSetCreditsFetchedActionCreator(true));
    }
}

export const applyCreditThunkCreator = () => {
    return async (dispatch) => {
        const data = await userAPI.applyCredit();
        return data;
    }
}
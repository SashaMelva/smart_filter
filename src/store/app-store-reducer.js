import {initialState} from "./initialState";
import {SET_ME, SET_CREDITS, SET_ME_FETCHED, SET_CREDITS_FETCHED} from "./action_types";

const appPageReducer = (state = initialState.appStore, action) => {
    let stateCopy = {...state};
    switch (action.type) {
        case SET_ME:
            stateCopy.me = action.data;
            return stateCopy;
            break;
        case SET_CREDITS:
            stateCopy.credits = action.data;
            return stateCopy;
            break;
        case SET_ME_FETCHED:
            stateCopy.meFetched = action.data;
            return stateCopy;
            break;
        case SET_CREDITS_FETCHED:
            stateCopy.creditsFetched = action.data;
            return stateCopy;
            break;
        default:
            return state;
    }
}

export default appPageReducer;
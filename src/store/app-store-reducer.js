import {initialState} from "./initialState";
import {SET_ME} from "./action_types";

const appPageReducer = (state = initialState.appStore, action) => {
    let stateCopy = {...state};
    switch (action.type) {
        case SET_ME:
            stateCopy.me = action.data;
            return stateCopy;
            break;
        default:
            return state;
    }
}

export default appPageReducer;

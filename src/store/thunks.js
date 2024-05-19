import {
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

import {applyMiddleware, combineReducers, createStore} from "redux";
import appStoreReducer from "./app-store-reducer";
import {thunk} from 'redux-thunk';

let reducers = combineReducers({
    appStore: appStoreReducer
});

let store = createStore(reducers, applyMiddleware(thunk));

export default store;
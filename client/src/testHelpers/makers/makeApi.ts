import {ApiClient} from "../../api/ApiClient";
import {Made} from "./made";

export const makeApi = (overrides: Partial<ApiClient> = {}): Made<ApiClient> =>
    Object.assign({
        authenticate: jest.fn(),
    }, overrides)

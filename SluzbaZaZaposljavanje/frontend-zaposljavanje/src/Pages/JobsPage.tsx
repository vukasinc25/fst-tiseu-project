import * as React from 'react';

function JobsPage() {
    return (
        <>
            <div className="row">
                <div className="col-sm-6">
                    <div className="card">
                        <div className="card-body">
                            <h5 className="card-title">Job #1</h5>
                            <p className="card-text">Bbbbbbb</p>
                            <a href="#" className="btn btn-primary">aaa</a>
                        </div>
                    </div>
                </div>
                <div className="col-sm-6">
                    <div className="card">
                        <div className="card-body">
                            <h5 className="card-title">Job #2</h5>
                            <p className="card-text">Aaaaaaaa</p>
                            <a href="#" className="btn btn-primary">bbb</a>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default JobsPage;

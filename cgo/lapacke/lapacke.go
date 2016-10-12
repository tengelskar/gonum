// Do not manually edit this file. It was created by the generate_lapacke.go from lapacke.h.

// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lapacke provides bindings to the LAPACKE C Interface to LAPACK.
//
// Links are provided to the NETLIB fortran implementation/dependencies for each function.
package lapacke

/*
#cgo CFLAGS: -g -O2
#include "lapacke.h"
*/
import "C"

import (
	"unsafe"

	"github.com/gonum/blas"
	"github.com/gonum/lapack"
)

// Type order is used to specify the matrix storage format. We still interact with
// an API that allows client calls to specify order, so this is here to document that fact.
type order int

const (
	rowMajor order = 101 + iota
	colMajor
)

func isZero(ret C.int) bool { return ret == 0 }

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sbdsdc.f.
func Sbdsdc(ul blas.Uplo, compq lapack.Comp, n int, d, e, u []float32, ldu int, vt []float32, ldvt int, q []float32, iq []int32, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float32
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _iq *int32
	if len(iq) > 0 {
		_iq = &iq[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sbdsdc_work((C.int)(rowMajor), (C.char)(ul), (C.char)(compq), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_vt), (C.lapack_int)(ldvt), (*C.float)(_q), (*C.lapack_int)(_iq), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dbdsdc.f.
func Dbdsdc(ul blas.Uplo, compq lapack.Comp, n int, d, e, u []float64, ldu int, vt []float64, ldvt int, q []float64, iq []int32, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _iq *int32
	if len(iq) > 0 {
		_iq = &iq[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dbdsdc_work((C.int)(rowMajor), (C.char)(ul), (C.char)(compq), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_vt), (C.lapack_int)(ldvt), (*C.double)(_q), (*C.lapack_int)(_iq), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sbdsvdx.f.
func Sbdsvdx(ul blas.Uplo, jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu, il, iu, ns int, s, z []float32, ldz int, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sbdsvdx_work((C.int)(rowMajor), (C.char)(ul), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.lapack_int)(vl), (C.lapack_int)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.lapack_int)(ns), (*C.float)(_s), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dbdsvdx.f.
func Dbdsvdx(ul blas.Uplo, jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu, il, iu, ns int, s, z []float64, ldz int, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dbdsvdx_work((C.int)(rowMajor), (C.char)(ul), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.lapack_int)(vl), (C.lapack_int)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.lapack_int)(ns), (*C.double)(_s), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sbdsqr.f.
func Sbdsqr(ul blas.Uplo, n, ncvt, nru, ncc int, d, e, vt []float32, ldvt int, u []float32, ldu int, c []float32, ldc int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _vt *float32
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sbdsqr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ncvt), (C.lapack_int)(nru), (C.lapack_int)(ncc), (*C.float)(_d), (*C.float)(_e), (*C.float)(_vt), (C.lapack_int)(ldvt), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dbdsqr.f.
func Dbdsqr(ul blas.Uplo, n, ncvt, nru, ncc int, d, e, vt []float64, ldvt int, u []float64, ldu int, c []float64, ldc int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _vt *float64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dbdsqr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ncvt), (C.lapack_int)(nru), (C.lapack_int)(ncc), (*C.double)(_d), (*C.double)(_e), (*C.double)(_vt), (C.lapack_int)(ldvt), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cbdsqr.f.
func Cbdsqr(ul blas.Uplo, n, ncvt, nru, ncc int, d, e []float32, vt []complex64, ldvt int, u []complex64, ldu int, c []complex64, ldc int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _vt *complex64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cbdsqr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ncvt), (C.lapack_int)(nru), (C.lapack_int)(ncc), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zbdsqr.f.
func Zbdsqr(ul blas.Uplo, n, ncvt, nru, ncc int, d, e []float64, vt []complex128, ldvt int, u []complex128, ldu int, c []complex128, ldc int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _vt *complex128
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zbdsqr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ncvt), (C.lapack_int)(nru), (C.lapack_int)(ncc), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sdisna.f.
func Sdisna(job lapack.Job, m, n int, d, sep []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _sep *float32
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	return isZero(C.LAPACKE_sdisna_work((C.char)(job), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_sep)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ddisna.f.
func Ddisna(job lapack.Job, m, n int, d, sep []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _sep *float64
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	return isZero(C.LAPACKE_ddisna_work((C.char)(job), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_sep)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbbrd.f.
func Sgbbrd(vect byte, m, n, ncc, kl, ku int, ab []float32, ldab int, d, e, q []float32, ldq int, pt []float32, ldpt int, c []float32, ldc int, work []float32) bool {
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _pt *float32
	if len(pt) > 0 {
		_pt = &pt[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgbbrd_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ncc), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_d), (*C.float)(_e), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_pt), (C.lapack_int)(ldpt), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbbrd.f.
func Dgbbrd(vect byte, m, n, ncc, kl, ku int, ab []float64, ldab int, d, e, q []float64, ldq int, pt []float64, ldpt int, c []float64, ldc int, work []float64) bool {
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _pt *float64
	if len(pt) > 0 {
		_pt = &pt[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgbbrd_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ncc), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_d), (*C.double)(_e), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_pt), (C.lapack_int)(ldpt), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbbrd.f.
func Cgbbrd(vect byte, m, n, ncc, kl, ku int, ab []complex64, ldab int, d, e []float32, q []complex64, ldq int, pt []complex64, ldpt int, c []complex64, ldc int, work []complex64, rwork []float32) bool {
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _pt *complex64
	if len(pt) > 0 {
		_pt = &pt[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgbbrd_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ncc), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_pt), (C.lapack_int)(ldpt), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbbrd.f.
func Zgbbrd(vect byte, m, n, ncc, kl, ku int, ab []complex128, ldab int, d, e []float64, q []complex128, ldq int, pt []complex128, ldpt int, c []complex128, ldc int, work []complex128, rwork []float64) bool {
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _pt *complex128
	if len(pt) > 0 {
		_pt = &pt[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgbbrd_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ncc), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_pt), (C.lapack_int)(ldpt), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbcon.f.
func Sgbcon(norm byte, n, kl, ku int, ab []float32, ldab int, ipiv []int32, anorm float32, rcond, work []float32, iwork []int32) bool {
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgbcon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbcon.f.
func Dgbcon(norm byte, n, kl, ku int, ab []float64, ldab int, ipiv []int32, anorm float64, rcond, work []float64, iwork []int32) bool {
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgbcon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbcon.f.
func Cgbcon(norm byte, n, kl, ku int, ab []complex64, ldab int, ipiv []int32, anorm float32, rcond []float32, work []complex64, rwork []float32) bool {
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgbcon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbcon.f.
func Zgbcon(norm byte, n, kl, ku int, ab []complex128, ldab int, ipiv []int32, anorm float64, rcond []float64, work []complex128, rwork []float64) bool {
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgbcon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbequ.f.
func Sgbequ(m, n, kl, ku int, ab []float32, ldab int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_sgbequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbequ.f.
func Dgbequ(m, n, kl, ku int, ab []float64, ldab int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dgbequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbequ.f.
func Cgbequ(m, n, kl, ku int, ab []complex64, ldab int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cgbequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbequ.f.
func Zgbequ(m, n, kl, ku int, ab []complex128, ldab int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zgbequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbequb.f.
func Sgbequb(m, n, kl, ku int, ab []float32, ldab int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_sgbequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbequb.f.
func Dgbequb(m, n, kl, ku int, ab []float64, ldab int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dgbequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbequb.f.
func Cgbequb(m, n, kl, ku int, ab []complex64, ldab int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cgbequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbequb.f.
func Zgbequb(m, n, kl, ku int, ab []complex128, ldab int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zgbequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbrfs.f.
func Sgbrfs(trans blas.Transpose, n, kl, ku, nrhs int, ab []float32, ldab int, afb []float32, ldafb int, ipiv []int32, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float32
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgbrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbrfs.f.
func Dgbrfs(trans blas.Transpose, n, kl, ku, nrhs int, ab []float64, ldab int, afb []float64, ldafb int, ipiv []int32, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgbrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbrfs.f.
func Cgbrfs(trans blas.Transpose, n, kl, ku, nrhs int, ab []complex64, ldab int, afb []complex64, ldafb int, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgbrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbrfs.f.
func Zgbrfs(trans blas.Transpose, n, kl, ku, nrhs int, ab []complex128, ldab int, afb []complex128, ldafb int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex128
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgbrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbsv.f.
func Sgbsv(n, kl, ku, nrhs int, ab []float32, ldab int, ipiv []int32, b []float32, ldb int) bool {
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sgbsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbsv.f.
func Dgbsv(n, kl, ku, nrhs int, ab []float64, ldab int, ipiv []int32, b []float64, ldb int) bool {
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dgbsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbsv.f.
func Cgbsv(n, kl, ku, nrhs int, ab []complex64, ldab int, ipiv []int32, b []complex64, ldb int) bool {
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cgbsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbsv.f.
func Zgbsv(n, kl, ku, nrhs int, ab []complex128, ldab int, ipiv []int32, b []complex128, ldb int) bool {
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zgbsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbsvx.f.
func Sgbsvx(fact byte, trans blas.Transpose, n, kl, ku, nrhs int, ab []float32, ldab int, afb []float32, ldafb int, ipiv []int32, equed []byte, r, c, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float32
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_r), (*C.float)(_c), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbsvx.f.
func Dgbsvx(fact byte, trans blas.Transpose, n, kl, ku, nrhs int, ab []float64, ldab int, afb []float64, ldafb int, ipiv []int32, equed []byte, r, c, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_r), (*C.double)(_c), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbsvx.f.
func Cgbsvx(fact byte, trans blas.Transpose, n, kl, ku, nrhs int, ab []complex64, ldab int, afb []complex64, ldafb int, ipiv []int32, equed []byte, r, c []float32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_r), (*C.float)(_c), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbsvx.f.
func Zgbsvx(fact byte, trans blas.Transpose, n, kl, ku, nrhs int, ab []complex128, ldab int, afb []complex128, ldafb int, ipiv []int32, equed []byte, r, c []float64, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex128
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_afb), (C.lapack_int)(ldafb), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_r), (*C.double)(_c), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbtrf.f.
func Sgbtrf(m, n, kl, ku int, ab []float32, ldab int, ipiv []int32) bool {
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_sgbtrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbtrf.f.
func Dgbtrf(m, n, kl, ku int, ab []float64, ldab int, ipiv []int32) bool {
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dgbtrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbtrf.f.
func Cgbtrf(m, n, kl, ku int, ab []complex64, ldab int, ipiv []int32) bool {
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_cgbtrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbtrf.f.
func Zgbtrf(m, n, kl, ku int, ab []complex128, ldab int, ipiv []int32) bool {
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zgbtrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgbtrs.f.
func Sgbtrs(trans blas.Transpose, n, kl, ku, nrhs int, ab []float32, ldab int, ipiv []int32, b []float32, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sgbtrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgbtrs.f.
func Dgbtrs(trans blas.Transpose, n, kl, ku, nrhs int, ab []float64, ldab int, ipiv []int32, b []float64, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dgbtrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgbtrs.f.
func Cgbtrs(trans blas.Transpose, n, kl, ku, nrhs int, ab []complex64, ldab int, ipiv []int32, b []complex64, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cgbtrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgbtrs.f.
func Zgbtrs(trans blas.Transpose, n, kl, ku, nrhs int, ab []complex128, ldab int, ipiv []int32, b []complex128, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zgbtrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgebak.f.
func Sgebak(job lapack.Job, side blas.Side, n, ilo, ihi int, scale []float32, m int, v []float32, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_sgebak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_scale), (C.lapack_int)(m), (*C.float)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgebak.f.
func Dgebak(job lapack.Job, side blas.Side, n, ilo, ihi int, scale []float64, m int, v []float64, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_dgebak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_scale), (C.lapack_int)(m), (*C.double)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgebak.f.
func Cgebak(job lapack.Job, side blas.Side, n, ilo, ihi int, scale []float32, m int, v []complex64, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_cgebak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_scale), (C.lapack_int)(m), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgebak.f.
func Zgebak(job lapack.Job, side blas.Side, n, ilo, ihi int, scale []float64, m int, v []complex128, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_zgebak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_scale), (C.lapack_int)(m), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgebal.f.
func Sgebal(job lapack.Job, n int, a []float32, lda int, ilo, ihi []int32, scale []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_sgebal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgebal.f.
func Dgebal(job lapack.Job, n int, a []float64, lda int, ilo, ihi []int32, scale []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_dgebal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgebal.f.
func Cgebal(job lapack.Job, n int, a []complex64, lda int, ilo, ihi []int32, scale []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_cgebal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgebal.f.
func Zgebal(job lapack.Job, n int, a []complex128, lda int, ilo, ihi []int32, scale []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_zgebal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgebrd.f.
func Sgebrd(m, n int, a []float32, lda int, d, e, tauq, taup, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tauq *float32
	if len(tauq) > 0 {
		_tauq = &tauq[0]
	}
	var _taup *float32
	if len(taup) > 0 {
		_taup = &taup[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgebrd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_d), (*C.float)(_e), (*C.float)(_tauq), (*C.float)(_taup), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgebrd.f.
func Dgebrd(m, n int, a []float64, lda int, d, e, tauq, taup, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tauq *float64
	if len(tauq) > 0 {
		_tauq = &tauq[0]
	}
	var _taup *float64
	if len(taup) > 0 {
		_taup = &taup[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgebrd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_d), (*C.double)(_e), (*C.double)(_tauq), (*C.double)(_taup), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgebrd.f.
func Cgebrd(m, n int, a []complex64, lda int, d, e []float32, tauq, taup, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tauq *complex64
	if len(tauq) > 0 {
		_tauq = &tauq[0]
	}
	var _taup *complex64
	if len(taup) > 0 {
		_taup = &taup[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgebrd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_tauq), (*C.lapack_complex_float)(_taup), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgebrd.f.
func Zgebrd(m, n int, a []complex128, lda int, d, e []float64, tauq, taup, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tauq *complex128
	if len(tauq) > 0 {
		_tauq = &tauq[0]
	}
	var _taup *complex128
	if len(taup) > 0 {
		_taup = &taup[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgebrd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_tauq), (*C.lapack_complex_double)(_taup), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgecon.f.
func Sgecon(norm byte, n int, a []float32, lda int, anorm float32, rcond, work []float32, iwork []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgecon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgecon.f.
func Dgecon(norm byte, n int, a []float64, lda int, anorm float64, rcond, work []float64, iwork []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgecon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgecon.f.
func Cgecon(norm byte, n int, a []complex64, lda int, anorm float32, rcond []float32, work []complex64, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgecon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgecon.f.
func Zgecon(norm byte, n int, a []complex128, lda int, anorm float64, rcond []float64, work []complex128, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgecon_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeequ.f.
func Sgeequ(m, n int, a []float32, lda int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_sgeequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeequ.f.
func Dgeequ(m, n int, a []float64, lda int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dgeequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeequ.f.
func Cgeequ(m, n int, a []complex64, lda int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cgeequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeequ.f.
func Zgeequ(m, n int, a []complex128, lda int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zgeequ_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeequb.f.
func Sgeequb(m, n int, a []float32, lda int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_sgeequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeequb.f.
func Dgeequb(m, n int, a []float64, lda int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dgeequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeequb.f.
func Cgeequb(m, n int, a []complex64, lda int, r, c, rowcnd, colcnd, amax []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float32
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float32
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cgeequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_r), (*C.float)(_c), (*C.float)(_rowcnd), (*C.float)(_colcnd), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeequb.f.
func Zgeequb(m, n int, a []complex128, lda int, r, c, rowcnd, colcnd, amax []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _rowcnd *float64
	if len(rowcnd) > 0 {
		_rowcnd = &rowcnd[0]
	}
	var _colcnd *float64
	if len(colcnd) > 0 {
		_colcnd = &colcnd[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zgeequb_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_r), (*C.double)(_c), (*C.double)(_rowcnd), (*C.double)(_colcnd), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeev.f.
func Sgeev(jobvl, jobvr lapack.Job, n int, a []float32, lda int, wr, wi, vl []float32, ldvl int, vr []float32, ldvr int, work []float32, lwork int) int {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _wr *float32
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float32
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return int(C.LAPACKE_sgeev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_wr), (*C.float)(_wi), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeev.f.
func Dgeev(jobvl, jobvr lapack.Job, n int, a []float64, lda int, wr, wi, vl []float64, ldvl int, vr []float64, ldvr int, work []float64, lwork int) int {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _wr *float64
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float64
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return int(C.LAPACKE_dgeev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_wr), (*C.double)(_wi), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeev.f.
func Cgeev(jobvl, jobvr lapack.Job, n int, a []complex64, lda int, w, vl []complex64, ldvl int, vr []complex64, ldvr int, work []complex64, lwork int, rwork []float32) int {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *complex64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return int(C.LAPACKE_cgeev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_w), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeev.f.
func Zgeev(jobvl, jobvr lapack.Job, n int, a []complex128, lda int, w, vl []complex128, ldvl int, vr []complex128, ldvr int, work []complex128, lwork int, rwork []float64) int {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *complex128
	if len(w) > 0 {
		_w = &w[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return int(C.LAPACKE_zgeev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_w), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeevx.f.
func Sgeevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []float32, lda int, wr, wi, vl []float32, ldvl int, vr []float32, ldvr int, ilo, ihi []int32, scale, abnrm, rconde, rcondv, work []float32, lwork int, iwork []int32) int {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _wr *float32
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float32
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _abnrm *float32
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _rconde *float32
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float32
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return int(C.LAPACKE_sgeevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_wr), (*C.float)(_wi), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_scale), (*C.float)(_abnrm), (*C.float)(_rconde), (*C.float)(_rcondv), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeevx.f.
func Dgeevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []float64, lda int, wr, wi, vl []float64, ldvl int, vr []float64, ldvr int, ilo, ihi []int32, scale, abnrm, rconde, rcondv, work []float64, lwork int, iwork []int32) int {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _wr *float64
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float64
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _abnrm *float64
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _rconde *float64
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float64
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return int(C.LAPACKE_dgeevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_wr), (*C.double)(_wi), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_scale), (*C.double)(_abnrm), (*C.double)(_rconde), (*C.double)(_rcondv), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeevx.f.
func Cgeevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []complex64, lda int, w, vl []complex64, ldvl int, vr []complex64, ldvr int, ilo, ihi []int32, scale, abnrm, rconde, rcondv []float32, work []complex64, lwork int, rwork []float32) int {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *complex64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _abnrm *float32
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _rconde *float32
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float32
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return int(C.LAPACKE_cgeevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_w), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_scale), (*C.float)(_abnrm), (*C.float)(_rconde), (*C.float)(_rcondv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeevx.f.
func Zgeevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []complex128, lda int, w, vl []complex128, ldvl int, vr []complex128, ldvr int, ilo, ihi []int32, scale, abnrm, rconde, rcondv []float64, work []complex128, lwork int, rwork []float64) int {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *complex128
	if len(w) > 0 {
		_w = &w[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _abnrm *float64
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _rconde *float64
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float64
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return int(C.LAPACKE_zgeevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_w), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_scale), (*C.double)(_abnrm), (*C.double)(_rconde), (*C.double)(_rcondv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgehrd.f.
func Sgehrd(n, ilo, ihi int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgehrd_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgehrd.f.
func Dgehrd(n, ilo, ihi int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgehrd_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgehrd.f.
func Cgehrd(n, ilo, ihi int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgehrd_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgehrd.f.
func Zgehrd(n, ilo, ihi int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgehrd_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgejsv.f.
func Sgejsv(joba, jobu, jobv, jobr, jobt, jobp lapack.Job, m, n int, a []float32, lda int, sva, u []float32, ldu int, v []float32, ldv int, work []float32, lwork int, iwork []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float32
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgejsv_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.char)(jobr), (C.char)(jobt), (C.char)(jobp), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_sva), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgejsv.f.
func Dgejsv(joba, jobu, jobv, jobr, jobt, jobp lapack.Job, m, n int, a []float64, lda int, sva, u []float64, ldu int, v []float64, ldv int, work []float64, lwork int, iwork []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float64
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgejsv_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.char)(jobr), (C.char)(jobt), (C.char)(jobp), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_sva), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgejsv.f.
func Cgejsv(joba, jobu, jobv, jobr, jobt, jobp lapack.Job, m, n int, a []complex64, lda int, sva []float32, u []complex64, ldu int, v []complex64, ldv int, cwork []complex64, lwork int, work []float32, lrwork int, iwork []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float32
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _cwork *complex64
	if len(cwork) > 0 {
		_cwork = &cwork[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cgejsv_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.char)(jobr), (C.char)(jobt), (C.char)(jobp), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_sva), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_cwork), (C.lapack_int)(lwork), (*C.float)(_work), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgejsv.f.
func Zgejsv(joba, jobu, jobv, jobr, jobt, jobp lapack.Job, m, n int, a []complex128, lda int, sva []float64, u []complex128, ldu int, v []complex128, ldv int, cwork []complex128, lwork int, work []float64, lrwork int, iwork []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float64
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _cwork *complex128
	if len(cwork) > 0 {
		_cwork = &cwork[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zgejsv_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.char)(jobr), (C.char)(jobt), (C.char)(jobp), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_sva), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_cwork), (C.lapack_int)(lwork), (*C.double)(_work), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgelq2.f.
func Sgelq2(m, n int, a []float32, lda int, tau, work []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgelq2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgelq2.f.
func Dgelq2(m, n int, a []float64, lda int, tau, work []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgelq2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgelq2.f.
func Cgelq2(m, n int, a []complex64, lda int, tau, work []complex64) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgelq2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgelq2.f.
func Zgelq2(m, n int, a []complex128, lda int, tau, work []complex128) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgelq2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgelqf.f.
func Sgelqf(m, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgelqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgelqf.f.
func Dgelqf(m, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgelqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgelqf.f.
func Cgelqf(m, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgelqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgelqf.f.
func Zgelqf(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgelqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgels.f.
func Sgels(trans blas.Transpose, m, n, nrhs int, a []float32, lda int, b []float32, ldb int, work []float32, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgels_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgels.f.
func Dgels(trans blas.Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int, work []float64, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgels_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgels.f.
func Cgels(trans blas.Transpose, m, n, nrhs int, a []complex64, lda int, b []complex64, ldb int, work []complex64, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgels_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgels.f.
func Zgels(trans blas.Transpose, m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, work []complex128, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgels_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgelsd.f.
func Sgelsd(m, n, nrhs int, a []float32, lda int, b []float32, ldb int, s []float32, rcond float32, rank []int32, work []float32, lwork int, iwork []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgelsd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_s), (C.float)(rcond), (*C.lapack_int)(_rank), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgelsd.f.
func Dgelsd(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, s []float64, rcond float64, rank []int32, work []float64, lwork int, iwork []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgelsd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_s), (C.double)(rcond), (*C.lapack_int)(_rank), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgelsd.f.
func Cgelsd(m, n, nrhs int, a []complex64, lda int, b []complex64, ldb int, s []float32, rcond float32, rank []int32, work []complex64, lwork int, rwork []float32, iwork []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cgelsd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.float)(_s), (C.float)(rcond), (*C.lapack_int)(_rank), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgelsd.f.
func Zgelsd(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, s []float64, rcond float64, rank []int32, work []complex128, lwork int, rwork []float64, iwork []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zgelsd_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.double)(_s), (C.double)(rcond), (*C.lapack_int)(_rank), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgelss.f.
func Sgelss(m, n, nrhs int, a []float32, lda int, b []float32, ldb int, s []float32, rcond float32, rank []int32, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgelss_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_s), (C.float)(rcond), (*C.lapack_int)(_rank), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgelss.f.
func Dgelss(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, s []float64, rcond float64, rank []int32, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgelss_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_s), (C.double)(rcond), (*C.lapack_int)(_rank), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgelss.f.
func Cgelss(m, n, nrhs int, a []complex64, lda int, b []complex64, ldb int, s []float32, rcond float32, rank []int32, work []complex64, lwork int, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgelss_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.float)(_s), (C.float)(rcond), (*C.lapack_int)(_rank), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgelss.f.
func Zgelss(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, s []float64, rcond float64, rank []int32, work []complex128, lwork int, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgelss_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.double)(_s), (C.double)(rcond), (*C.lapack_int)(_rank), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgelsy.f.
func Sgelsy(m, n, nrhs int, a []float32, lda int, b []float32, ldb int, jpvt []int32, rcond float32, rank []int32, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgelsy_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_jpvt), (C.float)(rcond), (*C.lapack_int)(_rank), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgelsy.f.
func Dgelsy(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, jpvt []int32, rcond float64, rank []int32, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgelsy_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_jpvt), (C.double)(rcond), (*C.lapack_int)(_rank), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgelsy.f.
func Cgelsy(m, n, nrhs int, a []complex64, lda int, b []complex64, ldb int, jpvt []int32, rcond float32, rank []int32, work []complex64, lwork int, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgelsy_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_jpvt), (C.float)(rcond), (*C.lapack_int)(_rank), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgelsy.f.
func Zgelsy(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, jpvt []int32, rcond float64, rank []int32, work []complex128, lwork int, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgelsy_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_jpvt), (C.double)(rcond), (*C.lapack_int)(_rank), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqlf.f.
func Sgeqlf(m, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgeqlf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqlf.f.
func Dgeqlf(m, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgeqlf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqlf.f.
func Cgeqlf(m, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgeqlf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqlf.f.
func Zgeqlf(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgeqlf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqp3.f.
func Sgeqp3(m, n int, a []float32, lda int, jpvt []int32, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgeqp3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_jpvt), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqp3.f.
func Dgeqp3(m, n int, a []float64, lda int, jpvt []int32, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgeqp3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_jpvt), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqp3.f.
func Cgeqp3(m, n int, a []complex64, lda int, jpvt []int32, tau, work []complex64, lwork int, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgeqp3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_jpvt), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqp3.f.
func Zgeqp3(m, n int, a []complex128, lda int, jpvt []int32, tau, work []complex128, lwork int, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _jpvt *int32
	if len(jpvt) > 0 {
		_jpvt = &jpvt[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgeqp3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_jpvt), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqr2.f.
func Sgeqr2(m, n int, a []float32, lda int, tau, work []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgeqr2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqr2.f.
func Dgeqr2(m, n int, a []float64, lda int, tau, work []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgeqr2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqr2.f.
func Cgeqr2(m, n int, a []complex64, lda int, tau, work []complex64) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgeqr2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqr2.f.
func Zgeqr2(m, n int, a []complex128, lda int, tau, work []complex128) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgeqr2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqrf.f.
func Sgeqrf(m, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgeqrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqrf.f.
func Dgeqrf(m, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgeqrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqrf.f.
func Cgeqrf(m, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgeqrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqrf.f.
func Zgeqrf(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgeqrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqrfp.f.
func Sgeqrfp(m, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgeqrfp_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqrfp.f.
func Dgeqrfp(m, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgeqrfp_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqrfp.f.
func Cgeqrfp(m, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgeqrfp_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqrfp.f.
func Zgeqrfp(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgeqrfp_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgerfs.f.
func Sgerfs(trans blas.Transpose, n, nrhs int, a []float32, lda int, af []float32, ldaf int, ipiv []int32, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float32
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgerfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgerfs.f.
func Dgerfs(trans blas.Transpose, n, nrhs int, a []float64, lda int, af []float64, ldaf int, ipiv []int32, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgerfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgerfs.f.
func Cgerfs(trans blas.Transpose, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgerfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgerfs.f.
func Zgerfs(trans blas.Transpose, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgerfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgerqf.f.
func Sgerqf(m, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgerqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgerqf.f.
func Dgerqf(m, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgerqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgerqf.f.
func Cgerqf(m, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgerqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgerqf.f.
func Zgerqf(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgerqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgesdd.f.
func Sgesdd(jobz lapack.Job, m, n int, a []float32, lda int, s, u []float32, ldu int, vt []float32, ldvt int, work []float32, lwork int, iwork []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float32
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgesdd_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_vt), (C.lapack_int)(ldvt), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgesdd.f.
func Dgesdd(jobz lapack.Job, m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int, iwork []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgesdd_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_vt), (C.lapack_int)(ldvt), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgesdd.f.
func Cgesdd(jobz lapack.Job, m, n int, a []complex64, lda int, s []float32, u []complex64, ldu int, vt []complex64, ldvt int, work []complex64, lwork int, rwork []float32, iwork []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *complex64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cgesdd_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgesdd.f.
func Zgesdd(jobz lapack.Job, m, n int, a []complex128, lda int, s []float64, u []complex128, ldu int, vt []complex128, ldvt int, work []complex128, lwork int, rwork []float64, iwork []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *complex128
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zgesdd_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgesv.f.
func Sgesv(n, nrhs int, a []float32, lda int, ipiv []int32, b []float32, ldb int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sgesv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgesv.f.
func Dgesv(n, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dgesv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgesv.f.
func Cgesv(n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cgesv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgesv.f.
func Zgesv(n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zgesv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsgesv.f.
func Dsgesv(n, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int, x []float64, ldx int, work []float64, swork []float32, iter []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _swork *float32
	if len(swork) > 0 {
		_swork = &swork[0]
	}
	var _iter *int32
	if len(iter) > 0 {
		_iter = &iter[0]
	}
	return isZero(C.LAPACKE_dsgesv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_work), (*C.float)(_swork), (*C.lapack_int)(_iter)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zcgesv.f.
func Zcgesv(n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, work []complex128, swork []complex64, rwork []float64, iter []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _swork *complex64
	if len(swork) > 0 {
		_swork = &swork[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iter *int32
	if len(iter) > 0 {
		_iter = &iter[0]
	}
	return isZero(C.LAPACKE_zcgesv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.lapack_complex_double)(_work), (*C.lapack_complex_float)(_swork), (*C.double)(_rwork), (*C.lapack_int)(_iter)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgesvd.f.
func Sgesvd(jobu, jobvt lapack.Job, m, n int, a []float32, lda int, s, u []float32, ldu int, vt []float32, ldvt int, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float32
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgesvd_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_vt), (C.lapack_int)(ldvt), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgesvd.f.
func Dgesvd(jobu, jobvt lapack.Job, m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgesvd_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_vt), (C.lapack_int)(ldvt), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgesvd.f.
func Cgesvd(jobu, jobvt lapack.Job, m, n int, a []complex64, lda int, s []float32, u []complex64, ldu int, vt []complex64, ldvt int, work []complex64, lwork int, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *complex64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgesvd_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgesvd.f.
func Zgesvd(jobu, jobvt lapack.Job, m, n int, a []complex128, lda int, s []float64, u []complex128, ldu int, vt []complex128, ldvt int, work []complex128, lwork int, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *complex128
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgesvd_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgesvdx.f.
func Sgesvdx(jobu, jobvt lapack.Job, rng byte, m, n int, a []float32, lda, vl, vu, il, iu, ns int, s, u []float32, ldu int, vt []float32, ldvt int, work []float32, lwork int, iwork []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float32
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgesvdx_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.char)(rng), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (C.lapack_int)(vl), (C.lapack_int)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.lapack_int)(ns), (*C.float)(_s), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_vt), (C.lapack_int)(ldvt), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgesvdx.f.
func Dgesvdx(jobu, jobvt lapack.Job, rng byte, m, n int, a []float64, lda, vl, vu, il, iu, ns int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int, iwork []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *float64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgesvdx_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.char)(rng), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (C.lapack_int)(vl), (C.lapack_int)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.lapack_int)(ns), (*C.double)(_s), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_vt), (C.lapack_int)(ldvt), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgesvdx.f.
func Cgesvdx(jobu, jobvt lapack.Job, rng byte, m, n int, a []complex64, lda, vl, vu, il, iu, ns int, s []float32, u []complex64, ldu int, vt []complex64, ldvt int, work []complex64, lwork int, rwork []float32, iwork []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *complex64
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cgesvdx_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.char)(rng), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.lapack_int)(vl), (C.lapack_int)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.lapack_int)(ns), (*C.float)(_s), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgesvdx.f.
func Zgesvdx(jobu, jobvt lapack.Job, rng byte, m, n int, a []complex128, lda, vl, vu, il, iu, ns int, s []float64, u []complex128, ldu int, vt []complex128, ldvt int, work []complex128, lwork int, rwork []float64, iwork []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _vt *complex128
	if len(vt) > 0 {
		_vt = &vt[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zgesvdx_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobvt), (C.char)(rng), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.lapack_int)(vl), (C.lapack_int)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.lapack_int)(ns), (*C.double)(_s), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_vt), (C.lapack_int)(ldvt), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgesvj.f.
func Sgesvj(joba, jobu, jobv lapack.Job, m, n int, a []float32, lda int, sva []float32, mv int, v []float32, ldv int, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float32
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgesvj_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_sva), (C.lapack_int)(mv), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgesvj.f.
func Dgesvj(joba, jobu, jobv lapack.Job, m, n int, a []float64, lda int, sva []float64, mv int, v []float64, ldv int, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float64
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgesvj_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_sva), (C.lapack_int)(mv), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgesvj.f.
func Cgesvj(joba, jobu, jobv lapack.Job, m, n int, a []complex64, lda int, sva []float32, mv int, v []complex64, ldv int, cwork []complex64, lwork int, rwork []float32, lrwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float32
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _cwork *complex64
	if len(cwork) > 0 {
		_cwork = &cwork[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgesvj_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_sva), (C.lapack_int)(mv), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_cwork), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgesvj.f.
func Zgesvj(joba, jobu, jobv lapack.Job, m, n int, a []complex128, lda int, sva []float64, mv int, v []complex128, ldv int, cwork []complex128, lwork int, rwork []float64, lrwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sva *float64
	if len(sva) > 0 {
		_sva = &sva[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _cwork *complex128
	if len(cwork) > 0 {
		_cwork = &cwork[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgesvj_work((C.int)(rowMajor), (C.char)(joba), (C.char)(jobu), (C.char)(jobv), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_sva), (C.lapack_int)(mv), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_cwork), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgesvx.f.
func Sgesvx(fact byte, trans blas.Transpose, n, nrhs int, a []float32, lda int, af []float32, ldaf int, ipiv []int32, equed []byte, r, c, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float32
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgesvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_r), (*C.float)(_c), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgesvx.f.
func Dgesvx(fact byte, trans blas.Transpose, n, nrhs int, a []float64, lda int, af []float64, ldaf int, ipiv []int32, equed []byte, r, c, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgesvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_r), (*C.double)(_c), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgesvx.f.
func Cgesvx(fact byte, trans blas.Transpose, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, ipiv []int32, equed []byte, r, c []float32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgesvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_r), (*C.float)(_c), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgesvx.f.
func Zgesvx(fact byte, trans blas.Transpose, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, ipiv []int32, equed []byte, r, c []float64, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgesvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_r), (*C.double)(_c), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgetf2.f.
func Sgetf2(m, n int, a []float32, lda int, ipiv []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_sgetf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgetf2.f.
func Dgetf2(m, n int, a []float64, lda int, ipiv []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dgetf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgetf2.f.
func Cgetf2(m, n int, a []complex64, lda int, ipiv []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_cgetf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgetf2.f.
func Zgetf2(m, n int, a []complex128, lda int, ipiv []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zgetf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgetrf.f.
func Sgetrf(m, n int, a []float32, lda int, ipiv []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_sgetrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgetrf.f.
func Dgetrf(m, n int, a []float64, lda int, ipiv []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dgetrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgetrf.f.
func Cgetrf(m, n int, a []complex64, lda int, ipiv []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_cgetrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgetrf.f.
func Zgetrf(m, n int, a []complex128, lda int, ipiv []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zgetrf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgetrf2.f.
func Sgetrf2(m, n int, a []float32, lda int, ipiv []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_sgetrf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgetrf2.f.
func Dgetrf2(m, n int, a []float64, lda int, ipiv []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dgetrf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgetrf2.f.
func Cgetrf2(m, n int, a []complex64, lda int, ipiv []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_cgetrf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgetrf2.f.
func Zgetrf2(m, n int, a []complex128, lda int, ipiv []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zgetrf2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgetri.f.
func Sgetri(n int, a []float32, lda int, ipiv []int32, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgetri_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgetri.f.
func Dgetri(n int, a []float64, lda int, ipiv []int32, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgetri_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgetri.f.
func Cgetri(n int, a []complex64, lda int, ipiv []int32, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgetri_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgetri.f.
func Zgetri(n int, a []complex128, lda int, ipiv []int32, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgetri_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgetrs.f.
func Sgetrs(trans blas.Transpose, n, nrhs int, a []float32, lda int, ipiv []int32, b []float32, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sgetrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgetrs.f.
func Dgetrs(trans blas.Transpose, n, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dgetrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgetrs.f.
func Cgetrs(trans blas.Transpose, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cgetrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgetrs.f.
func Zgetrs(trans blas.Transpose, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zgetrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggbak.f.
func Sggbak(job lapack.Job, side blas.Side, n, ilo, ihi int, lscale, rscale []float32, m int, v []float32, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _lscale *float32
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float32
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_sggbak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_lscale), (*C.float)(_rscale), (C.lapack_int)(m), (*C.float)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggbak.f.
func Dggbak(job lapack.Job, side blas.Side, n, ilo, ihi int, lscale, rscale []float64, m int, v []float64, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _lscale *float64
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float64
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_dggbak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_lscale), (*C.double)(_rscale), (C.lapack_int)(m), (*C.double)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggbak.f.
func Cggbak(job lapack.Job, side blas.Side, n, ilo, ihi int, lscale, rscale []float32, m int, v []complex64, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _lscale *float32
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float32
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_cggbak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_lscale), (*C.float)(_rscale), (C.lapack_int)(m), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggbak.f.
func Zggbak(job lapack.Job, side blas.Side, n, ilo, ihi int, lscale, rscale []float64, m int, v []complex128, ldv int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _lscale *float64
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float64
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	return isZero(C.LAPACKE_zggbak_work((C.int)(rowMajor), (C.char)(job), (C.char)(side), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_lscale), (*C.double)(_rscale), (C.lapack_int)(m), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggbal.f.
func Sggbal(job lapack.Job, n int, a []float32, lda int, b []float32, ldb int, ilo, ihi []int32, lscale, rscale, work []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float32
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float32
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggbal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_lscale), (*C.float)(_rscale), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggbal.f.
func Dggbal(job lapack.Job, n int, a []float64, lda int, b []float64, ldb int, ilo, ihi []int32, lscale, rscale, work []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float64
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float64
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggbal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_lscale), (*C.double)(_rscale), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggbal.f.
func Cggbal(job lapack.Job, n int, a []complex64, lda int, b []complex64, ldb int, ilo, ihi []int32, lscale, rscale, work []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float32
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float32
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cggbal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_lscale), (*C.float)(_rscale), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggbal.f.
func Zggbal(job lapack.Job, n int, a []complex128, lda int, b []complex128, ldb int, ilo, ihi []int32, lscale, rscale, work []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float64
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float64
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zggbal_work((C.int)(rowMajor), (C.char)(job), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_lscale), (*C.double)(_rscale), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggev.f.
func Sggev(jobvl, jobvr lapack.Job, n int, a []float32, lda int, b []float32, ldb int, alphar, alphai, beta, vl []float32, ldvl int, vr []float32, ldvr int, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float32
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float32
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_alphar), (*C.float)(_alphai), (*C.float)(_beta), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggev.f.
func Dggev(jobvl, jobvr lapack.Job, n int, a []float64, lda int, b []float64, ldb int, alphar, alphai, beta, vl []float64, ldvl int, vr []float64, ldvr int, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float64
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float64
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_alphar), (*C.double)(_alphai), (*C.double)(_beta), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggev.f.
func Cggev(jobvl, jobvr lapack.Job, n int, a []complex64, lda int, b []complex64, ldb int, alpha, beta, vl []complex64, ldvl int, vr []complex64, ldvr int, work []complex64, lwork int, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cggev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_alpha), (*C.lapack_complex_float)(_beta), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggev.f.
func Zggev(jobvl, jobvr lapack.Job, n int, a []complex128, lda int, b []complex128, ldb int, alpha, beta, vl []complex128, ldvl int, vr []complex128, ldvr int, work []complex128, lwork int, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex128
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex128
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zggev_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_alpha), (*C.lapack_complex_double)(_beta), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggev3.f.
func Sggev3(jobvl, jobvr lapack.Job, n int, a []float32, lda int, b []float32, ldb int, alphar, alphai, beta, vl []float32, ldvl int, vr []float32, ldvr int, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float32
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float32
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggev3_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_alphar), (*C.float)(_alphai), (*C.float)(_beta), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggev3.f.
func Dggev3(jobvl, jobvr lapack.Job, n int, a []float64, lda int, b []float64, ldb int, alphar, alphai, beta, vl []float64, ldvl int, vr []float64, ldvr int, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float64
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float64
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggev3_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_alphar), (*C.double)(_alphai), (*C.double)(_beta), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggev3.f.
func Cggev3(jobvl, jobvr lapack.Job, n int, a []complex64, lda int, b []complex64, ldb int, alpha, beta, vl []complex64, ldvl int, vr []complex64, ldvr int, work []complex64, lwork int, rwork []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cggev3_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_alpha), (*C.lapack_complex_float)(_beta), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggev3.f.
func Zggev3(jobvl, jobvr lapack.Job, n int, a []complex128, lda int, b []complex128, ldb int, alpha, beta, vl []complex128, ldvl int, vr []complex128, ldvr int, work []complex128, lwork int, rwork []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex128
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex128
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zggev3_work((C.int)(rowMajor), (C.char)(jobvl), (C.char)(jobvr), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_alpha), (*C.lapack_complex_double)(_beta), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggevx.f.
func Sggevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []float32, lda int, b []float32, ldb int, alphar, alphai, beta, vl []float32, ldvl int, vr []float32, ldvr int, ilo, ihi []int32, lscale, rscale, abnrm, bbnrm, rconde, rcondv, work []float32, lwork int, iwork, bwork []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float32
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float32
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float32
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float32
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _abnrm *float32
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _bbnrm *float32
	if len(bbnrm) > 0 {
		_bbnrm = &bbnrm[0]
	}
	var _rconde *float32
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float32
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _bwork *int32
	if len(bwork) > 0 {
		_bwork = &bwork[0]
	}
	return isZero(C.LAPACKE_sggevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_alphar), (*C.float)(_alphai), (*C.float)(_beta), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_lscale), (*C.float)(_rscale), (*C.float)(_abnrm), (*C.float)(_bbnrm), (*C.float)(_rconde), (*C.float)(_rcondv), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (*C.lapack_logical)(_bwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggevx.f.
func Dggevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []float64, lda int, b []float64, ldb int, alphar, alphai, beta, vl []float64, ldvl int, vr []float64, ldvr int, ilo, ihi []int32, lscale, rscale, abnrm, bbnrm, rconde, rcondv, work []float64, lwork int, iwork, bwork []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float64
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float64
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float64
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float64
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _abnrm *float64
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _bbnrm *float64
	if len(bbnrm) > 0 {
		_bbnrm = &bbnrm[0]
	}
	var _rconde *float64
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float64
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _bwork *int32
	if len(bwork) > 0 {
		_bwork = &bwork[0]
	}
	return isZero(C.LAPACKE_dggevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_alphar), (*C.double)(_alphai), (*C.double)(_beta), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_lscale), (*C.double)(_rscale), (*C.double)(_abnrm), (*C.double)(_bbnrm), (*C.double)(_rconde), (*C.double)(_rcondv), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (*C.lapack_logical)(_bwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggevx.f.
func Cggevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []complex64, lda int, b []complex64, ldb int, alpha, beta, vl []complex64, ldvl int, vr []complex64, ldvr int, ilo, ihi []int32, lscale, rscale, abnrm, bbnrm, rconde, rcondv []float32, work []complex64, lwork int, rwork []float32, iwork, bwork []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float32
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float32
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _abnrm *float32
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _bbnrm *float32
	if len(bbnrm) > 0 {
		_bbnrm = &bbnrm[0]
	}
	var _rconde *float32
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float32
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _bwork *int32
	if len(bwork) > 0 {
		_bwork = &bwork[0]
	}
	return isZero(C.LAPACKE_cggevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_alpha), (*C.lapack_complex_float)(_beta), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.float)(_lscale), (*C.float)(_rscale), (*C.float)(_abnrm), (*C.float)(_bbnrm), (*C.float)(_rconde), (*C.float)(_rcondv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_logical)(_bwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggevx.f.
func Zggevx(balanc byte, jobvl, jobvr lapack.Job, sense byte, n int, a []complex128, lda int, b []complex128, ldb int, alpha, beta, vl []complex128, ldvl int, vr []complex128, ldvr int, ilo, ihi []int32, lscale, rscale, abnrm, bbnrm, rconde, rcondv []float64, work []complex128, lwork int, rwork []float64, iwork, bwork []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex128
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex128
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _ilo *int32
	if len(ilo) > 0 {
		_ilo = &ilo[0]
	}
	var _ihi *int32
	if len(ihi) > 0 {
		_ihi = &ihi[0]
	}
	var _lscale *float64
	if len(lscale) > 0 {
		_lscale = &lscale[0]
	}
	var _rscale *float64
	if len(rscale) > 0 {
		_rscale = &rscale[0]
	}
	var _abnrm *float64
	if len(abnrm) > 0 {
		_abnrm = &abnrm[0]
	}
	var _bbnrm *float64
	if len(bbnrm) > 0 {
		_bbnrm = &bbnrm[0]
	}
	var _rconde *float64
	if len(rconde) > 0 {
		_rconde = &rconde[0]
	}
	var _rcondv *float64
	if len(rcondv) > 0 {
		_rcondv = &rcondv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _bwork *int32
	if len(bwork) > 0 {
		_bwork = &bwork[0]
	}
	return isZero(C.LAPACKE_zggevx_work((C.int)(rowMajor), (C.char)(balanc), (C.char)(jobvl), (C.char)(jobvr), (C.char)(sense), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_alpha), (*C.lapack_complex_double)(_beta), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.lapack_int)(_ilo), (*C.lapack_int)(_ihi), (*C.double)(_lscale), (*C.double)(_rscale), (*C.double)(_abnrm), (*C.double)(_bbnrm), (*C.double)(_rconde), (*C.double)(_rcondv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_logical)(_bwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggglm.f.
func Sggglm(n, m, p int, a []float32, lda int, b []float32, ldb int, d, x, y, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _y *float32
	if len(y) > 0 {
		_y = &y[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggglm_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_d), (*C.float)(_x), (*C.float)(_y), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggglm.f.
func Dggglm(n, m, p int, a []float64, lda int, b []float64, ldb int, d, x, y, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _y *float64
	if len(y) > 0 {
		_y = &y[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggglm_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_d), (*C.double)(_x), (*C.double)(_y), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggglm.f.
func Cggglm(n, m, p int, a []complex64, lda int, b []complex64, ldb int, d, x, y, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _y *complex64
	if len(y) > 0 {
		_y = &y[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cggglm_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_x), (*C.lapack_complex_float)(_y), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggglm.f.
func Zggglm(n, m, p int, a []complex128, lda int, b []complex128, ldb int, d, x, y, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _y *complex128
	if len(y) > 0 {
		_y = &y[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zggglm_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_x), (*C.lapack_complex_double)(_y), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgghrd.f.
func Sgghrd(compq, compz lapack.Comp, n, ilo, ihi int, a []float32, lda int, b []float32, ldb int, q []float32, ldq int, z []float32, ldz int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	return isZero(C.LAPACKE_sgghrd_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_z), (C.lapack_int)(ldz)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgghrd.f.
func Dgghrd(compq, compz lapack.Comp, n, ilo, ihi int, a []float64, lda int, b []float64, ldb int, q []float64, ldq int, z []float64, ldz int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	return isZero(C.LAPACKE_dgghrd_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_z), (C.lapack_int)(ldz)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgghrd.f.
func Cgghrd(compq, compz lapack.Comp, n, ilo, ihi int, a []complex64, lda int, b []complex64, ldb int, q []complex64, ldq int, z []complex64, ldz int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	return isZero(C.LAPACKE_cgghrd_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgghrd.f.
func Zgghrd(compq, compz lapack.Comp, n, ilo, ihi int, a []complex128, lda int, b []complex128, ldb int, q []complex128, ldq int, z []complex128, ldz int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	return isZero(C.LAPACKE_zgghrd_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgghd3.f.
func Sgghd3(compq, compz lapack.Comp, n, ilo, ihi int, a []float32, lda int, b []float32, ldb int, q []float32, ldq int, z []float32, ldz int, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgghd3_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgghd3.f.
func Dgghd3(compq, compz lapack.Comp, n, ilo, ihi int, a []float64, lda int, b []float64, ldb int, q []float64, ldq int, z []float64, ldz int, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgghd3_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgghd3.f.
func Cgghd3(compq, compz lapack.Comp, n, ilo, ihi int, a []complex64, lda int, b []complex64, ldb int, q []complex64, ldq int, z []complex64, ldz int, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgghd3_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgghd3.f.
func Zgghd3(compq, compz lapack.Comp, n, ilo, ihi int, a []complex128, lda int, b []complex128, ldb int, q []complex128, ldq int, z []complex128, ldz int, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgghd3_work((C.int)(rowMajor), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgglse.f.
func Sgglse(m, n, p int, a []float32, lda int, b []float32, ldb int, c, d, x, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgglse_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_c), (*C.float)(_d), (*C.float)(_x), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgglse.f.
func Dgglse(m, n, p int, a []float64, lda int, b []float64, ldb int, c, d, x, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgglse_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_c), (*C.double)(_d), (*C.double)(_x), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgglse.f.
func Cgglse(m, n, p int, a []complex64, lda int, b []complex64, ldb int, c, d, x, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgglse_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_c), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_x), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgglse.f.
func Zgglse(m, n, p int, a []complex128, lda int, b []complex128, ldb int, c, d, x, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgglse_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_c), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_x), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggqrf.f.
func Sggqrf(n, m, p int, a []float32, lda int, taua, b []float32, ldb int, taub, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *float32
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *float32
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggqrf_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_taua), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_taub), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggqrf.f.
func Dggqrf(n, m, p int, a []float64, lda int, taua, b []float64, ldb int, taub, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *float64
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *float64
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggqrf_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_taua), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_taub), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggqrf.f.
func Cggqrf(n, m, p int, a []complex64, lda int, taua, b []complex64, ldb int, taub, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *complex64
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *complex64
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cggqrf_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_taua), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_taub), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggqrf.f.
func Zggqrf(n, m, p int, a []complex128, lda int, taua, b []complex128, ldb int, taub, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *complex128
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *complex128
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zggqrf_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(m), (C.lapack_int)(p), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_taua), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_taub), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggrqf.f.
func Sggrqf(m, p, n int, a []float32, lda int, taua, b []float32, ldb int, taub, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *float32
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *float32
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggrqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_taua), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_taub), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggrqf.f.
func Dggrqf(m, p, n int, a []float64, lda int, taua, b []float64, ldb int, taub, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *float64
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *float64
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggrqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_taua), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_taub), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggrqf.f.
func Cggrqf(m, p, n int, a []complex64, lda int, taua, b []complex64, ldb int, taub, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *complex64
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *complex64
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cggrqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_taua), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_taub), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggrqf.f.
func Zggrqf(m, p, n int, a []complex128, lda int, taua, b []complex128, ldb int, taub, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _taua *complex128
	if len(taua) > 0 {
		_taua = &taua[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _taub *complex128
	if len(taub) > 0 {
		_taub = &taub[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zggrqf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_taua), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_taub), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggsvd3.f.
func Sggsvd3(jobu, jobv, jobq lapack.Job, m, n, p int, k, l []int32, a []float32, lda int, b []float32, ldb int, alpha, beta, u []float32, ldu int, v []float32, ldv int, q []float32, ldq int, work []float32, lwork int, iwork []int32) bool {
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float32
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sggsvd3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_alpha), (*C.float)(_beta), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggsvd3.f.
func Dggsvd3(jobu, jobv, jobq lapack.Job, m, n, p int, k, l []int32, a []float64, lda int, b []float64, ldb int, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64, lwork int, iwork []int32) bool {
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dggsvd3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_alpha), (*C.double)(_beta), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggsvd3.f.
func Cggsvd3(jobu, jobv, jobq lapack.Job, m, n, p int, k, l []int32, a []complex64, lda int, b []complex64, ldb int, alpha, beta []float32, u []complex64, ldu int, v []complex64, ldv int, q []complex64, ldq int, work []complex64, lwork int, rwork []float32, iwork []int32) bool {
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float32
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cggsvd3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.float)(_alpha), (*C.float)(_beta), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggsvd3.f.
func Zggsvd3(jobu, jobv, jobq lapack.Job, m, n, p int, k, l []int32, a []complex128, lda int, b []complex128, ldb int, alpha, beta []float64, u []complex128, ldu int, v []complex128, ldv int, q []complex128, ldq int, work []complex128, lwork int, rwork []float64, iwork []int32) bool {
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zggsvd3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(p), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.double)(_alpha), (*C.double)(_beta), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sggsvp3.f.
func Sggsvp3(jobu, jobv, jobq lapack.Job, m, p, n int, a []float32, lda int, b []float32, ldb int, tola, tolb float32, k, l []int32, u []float32, ldu int, v []float32, ldv int, q []float32, ldq int, iwork []int32, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sggsvp3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (C.float)(tola), (C.float)(tolb), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_q), (C.lapack_int)(ldq), (*C.lapack_int)(_iwork), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dggsvp3.f.
func Dggsvp3(jobu, jobv, jobq lapack.Job, m, p, n int, a []float64, lda int, b []float64, ldb int, tola, tolb float64, k, l []int32, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, iwork []int32, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dggsvp3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (C.double)(tola), (C.double)(tolb), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_q), (C.lapack_int)(ldq), (*C.lapack_int)(_iwork), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cggsvp3.f.
func Cggsvp3(jobu, jobv, jobq lapack.Job, m, p, n int, a []complex64, lda int, b []complex64, ldb int, tola, tolb float32, k, l []int32, u []complex64, ldu int, v []complex64, ldv int, q []complex64, ldq int, iwork []int32, rwork []float32, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cggsvp3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (C.float)(tola), (C.float)(tolb), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_int)(_iwork), (*C.float)(_rwork), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zggsvp3.f.
func Zggsvp3(jobu, jobv, jobq lapack.Job, m, p, n int, a []complex128, lda int, b []complex128, ldb int, tola, tolb float64, k, l []int32, u []complex128, ldu int, v []complex128, ldv int, q []complex128, ldq int, iwork []int32, rwork []float64, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	var _l *int32
	if len(l) > 0 {
		_l = &l[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zggsvp3_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (C.double)(tola), (C.double)(tolb), (*C.lapack_int)(_k), (*C.lapack_int)(_l), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_int)(_iwork), (*C.double)(_rwork), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgtcon.f.
func Sgtcon(norm byte, n int, dl, d, du, du2 []float32, ipiv []int32, anorm float32, rcond, work []float32, iwork []int32) bool {
	var _dl *float32
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float32
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *float32
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgtcon_work((C.char)(norm), (C.lapack_int)(n), (*C.float)(_dl), (*C.float)(_d), (*C.float)(_du), (*C.float)(_du2), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgtcon.f.
func Dgtcon(norm byte, n int, dl, d, du, du2 []float64, ipiv []int32, anorm float64, rcond, work []float64, iwork []int32) bool {
	var _dl *float64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *float64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgtcon_work((C.char)(norm), (C.lapack_int)(n), (*C.double)(_dl), (*C.double)(_d), (*C.double)(_du), (*C.double)(_du2), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgtcon.f.
func Cgtcon(norm byte, n int, dl, d, du, du2 []complex64, ipiv []int32, anorm float32, rcond []float32, work []complex64) bool {
	var _dl *complex64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *complex64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgtcon_work((C.char)(norm), (C.lapack_int)(n), (*C.lapack_complex_float)(_dl), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_du), (*C.lapack_complex_float)(_du2), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgtcon.f.
func Zgtcon(norm byte, n int, dl, d, du, du2 []complex128, ipiv []int32, anorm float64, rcond []float64, work []complex128) bool {
	var _dl *complex128
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex128
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *complex128
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgtcon_work((C.char)(norm), (C.lapack_int)(n), (*C.lapack_complex_double)(_dl), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_du), (*C.lapack_complex_double)(_du2), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgtrfs.f.
func Sgtrfs(trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []float32, ipiv []int32, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *float32
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float32
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *float32
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *float32
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *float32
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *float32
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgtrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_dl), (*C.float)(_d), (*C.float)(_du), (*C.float)(_dlf), (*C.float)(_df), (*C.float)(_duf), (*C.float)(_du2), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgtrfs.f.
func Dgtrfs(trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []float64, ipiv []int32, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *float64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *float64
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *float64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *float64
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *float64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgtrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_dl), (*C.double)(_d), (*C.double)(_du), (*C.double)(_dlf), (*C.double)(_df), (*C.double)(_duf), (*C.double)(_du2), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgtrfs.f.
func Cgtrfs(trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []complex64, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *complex64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *complex64
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *complex64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *complex64
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *complex64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgtrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_dl), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_du), (*C.lapack_complex_float)(_dlf), (*C.lapack_complex_float)(_df), (*C.lapack_complex_float)(_duf), (*C.lapack_complex_float)(_du2), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgtrfs.f.
func Zgtrfs(trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []complex128, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *complex128
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex128
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *complex128
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *complex128
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *complex128
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *complex128
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgtrfs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_dl), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_du), (*C.lapack_complex_double)(_dlf), (*C.lapack_complex_double)(_df), (*C.lapack_complex_double)(_duf), (*C.lapack_complex_double)(_du2), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgtsv.f.
func Sgtsv(n, nrhs int, dl, d, du, b []float32, ldb int) bool {
	var _dl *float32
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float32
	if len(du) > 0 {
		_du = &du[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sgtsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_dl), (*C.float)(_d), (*C.float)(_du), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgtsv.f.
func Dgtsv(n, nrhs int, dl, d, du, b []float64, ldb int) bool {
	var _dl *float64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dgtsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_dl), (*C.double)(_d), (*C.double)(_du), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgtsv.f.
func Cgtsv(n, nrhs int, dl, d, du, b []complex64, ldb int) bool {
	var _dl *complex64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cgtsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_dl), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_du), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgtsv.f.
func Zgtsv(n, nrhs int, dl, d, du, b []complex128, ldb int) bool {
	var _dl *complex128
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex128
	if len(du) > 0 {
		_du = &du[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zgtsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_dl), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_du), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgtsvx.f.
func Sgtsvx(fact byte, trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []float32, ipiv []int32, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *float32
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float32
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *float32
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *float32
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *float32
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *float32
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sgtsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_dl), (*C.float)(_d), (*C.float)(_du), (*C.float)(_dlf), (*C.float)(_df), (*C.float)(_duf), (*C.float)(_du2), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgtsvx.f.
func Dgtsvx(fact byte, trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []float64, ipiv []int32, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *float64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *float64
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *float64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *float64
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *float64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dgtsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_dl), (*C.double)(_d), (*C.double)(_du), (*C.double)(_dlf), (*C.double)(_df), (*C.double)(_duf), (*C.double)(_du2), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgtsvx.f.
func Cgtsvx(fact byte, trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []complex64, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *complex64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *complex64
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *complex64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *complex64
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *complex64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cgtsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_dl), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_du), (*C.lapack_complex_float)(_dlf), (*C.lapack_complex_float)(_df), (*C.lapack_complex_float)(_duf), (*C.lapack_complex_float)(_du2), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgtsvx.f.
func Zgtsvx(fact byte, trans blas.Transpose, n, nrhs int, dl, d, du, dlf, df, duf, du2 []complex128, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *complex128
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex128
	if len(du) > 0 {
		_du = &du[0]
	}
	var _dlf *complex128
	if len(dlf) > 0 {
		_dlf = &dlf[0]
	}
	var _df *complex128
	if len(df) > 0 {
		_df = &df[0]
	}
	var _duf *complex128
	if len(duf) > 0 {
		_duf = &duf[0]
	}
	var _du2 *complex128
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zgtsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_dl), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_du), (*C.lapack_complex_double)(_dlf), (*C.lapack_complex_double)(_df), (*C.lapack_complex_double)(_duf), (*C.lapack_complex_double)(_du2), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgttrf.f.
func Sgttrf(n int, dl, d, du, du2 []float32, ipiv []int32) bool {
	var _dl *float32
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float32
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *float32
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_sgttrf_work((C.lapack_int)(n), (*C.float)(_dl), (*C.float)(_d), (*C.float)(_du), (*C.float)(_du2), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgttrf.f.
func Dgttrf(n int, dl, d, du, du2 []float64, ipiv []int32) bool {
	var _dl *float64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *float64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dgttrf_work((C.lapack_int)(n), (*C.double)(_dl), (*C.double)(_d), (*C.double)(_du), (*C.double)(_du2), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgttrf.f.
func Cgttrf(n int, dl, d, du, du2 []complex64, ipiv []int32) bool {
	var _dl *complex64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *complex64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_cgttrf_work((C.lapack_int)(n), (*C.lapack_complex_float)(_dl), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_du), (*C.lapack_complex_float)(_du2), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgttrf.f.
func Zgttrf(n int, dl, d, du, du2 []complex128, ipiv []int32) bool {
	var _dl *complex128
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex128
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *complex128
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zgttrf_work((C.lapack_int)(n), (*C.lapack_complex_double)(_dl), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_du), (*C.lapack_complex_double)(_du2), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgttrs.f.
func Sgttrs(trans blas.Transpose, n, nrhs int, dl, d, du, du2 []float32, ipiv []int32, b []float32, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *float32
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float32
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *float32
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sgttrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_dl), (*C.float)(_d), (*C.float)(_du), (*C.float)(_du2), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgttrs.f.
func Dgttrs(trans blas.Transpose, n, nrhs int, dl, d, du, du2 []float64, ipiv []int32, b []float64, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *float64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *float64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *float64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dgttrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_dl), (*C.double)(_d), (*C.double)(_du), (*C.double)(_du2), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgttrs.f.
func Cgttrs(trans blas.Transpose, n, nrhs int, dl, d, du, du2 []complex64, ipiv []int32, b []complex64, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *complex64
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex64
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *complex64
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cgttrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_dl), (*C.lapack_complex_float)(_d), (*C.lapack_complex_float)(_du), (*C.lapack_complex_float)(_du2), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgttrs.f.
func Zgttrs(trans blas.Transpose, n, nrhs int, dl, d, du, du2 []complex128, ipiv []int32, b []complex128, ldb int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _dl *complex128
	if len(dl) > 0 {
		_dl = &dl[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _du *complex128
	if len(du) > 0 {
		_du = &du[0]
	}
	var _du2 *complex128
	if len(du2) > 0 {
		_du2 = &du2[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zgttrs_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_dl), (*C.lapack_complex_double)(_d), (*C.lapack_complex_double)(_du), (*C.lapack_complex_double)(_du2), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbev.f.
func Chbev(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []complex64, ldab int, w []float32, z []complex64, ldz int, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chbev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbev.f.
func Zhbev(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []complex128, ldab int, w []float64, z []complex128, ldz int, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhbev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbevd.f.
func Chbevd(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []complex64, ldab int, w []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_chbevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbevd.f.
func Zhbevd(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []complex128, ldab int, w []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zhbevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbevx.f.
func Chbevx(jobz lapack.Job, rng byte, ul blas.Uplo, n, kd int, ab []complex64, ldab int, q []complex64, ldq int, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, work []complex64, rwork []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_chbevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbevx.f.
func Zhbevx(jobz lapack.Job, rng byte, ul blas.Uplo, n, kd int, ab []complex128, ldab int, q []complex128, ldq int, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, work []complex128, rwork []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_zhbevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbgst.f.
func Chbgst(vect byte, ul blas.Uplo, n, ka, kb int, ab []complex64, ldab int, bb []complex64, ldbb int, x []complex64, ldx int, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chbgst_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_bb), (C.lapack_int)(ldbb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbgst.f.
func Zhbgst(vect byte, ul blas.Uplo, n, ka, kb int, ab []complex128, ldab int, bb []complex128, ldbb int, x []complex128, ldx int, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex128
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhbgst_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_bb), (C.lapack_int)(ldbb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbgv.f.
func Chbgv(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []complex64, ldab int, bb []complex64, ldbb int, w []float32, z []complex64, ldz int, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chbgv_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_bb), (C.lapack_int)(ldbb), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbgv.f.
func Zhbgv(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []complex128, ldab int, bb []complex128, ldbb int, w []float64, z []complex128, ldz int, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex128
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhbgv_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_bb), (C.lapack_int)(ldbb), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbgvd.f.
func Chbgvd(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []complex64, ldab int, bb []complex64, ldbb int, w []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_chbgvd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_bb), (C.lapack_int)(ldbb), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbgvd.f.
func Zhbgvd(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []complex128, ldab int, bb []complex128, ldbb int, w []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex128
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zhbgvd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_bb), (C.lapack_int)(ldbb), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbgvx.f.
func Chbgvx(jobz lapack.Job, rng byte, ul blas.Uplo, n, ka, kb int, ab []complex64, ldab int, bb []complex64, ldbb int, q []complex64, ldq int, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, work []complex64, rwork []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_chbgvx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_bb), (C.lapack_int)(ldbb), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbgvx.f.
func Zhbgvx(jobz lapack.Job, rng byte, ul blas.Uplo, n, ka, kb int, ab []complex128, ldab int, bb []complex128, ldbb int, q []complex128, ldq int, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, work []complex128, rwork []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *complex128
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_zhbgvx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_bb), (C.lapack_int)(ldbb), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chbtrd.f.
func Chbtrd(vect byte, ul blas.Uplo, n, kd int, ab []complex64, ldab int, d, e []float32, q []complex64, ldq int, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chbtrd_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhbtrd.f.
func Zhbtrd(vect byte, ul blas.Uplo, n, kd int, ab []complex128, ldab int, d, e []float64, q []complex128, ldq int, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhbtrd_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/checon.f.
func Checon(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, anorm float32, rcond []float32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_checon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhecon.f.
func Zhecon(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, anorm float64, rcond []float64, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhecon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cheequb.f.
func Cheequb(ul blas.Uplo, n int, a []complex64, lda int, s, scond, amax []float32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cheequb_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zheequb.f.
func Zheequb(ul blas.Uplo, n int, a []complex128, lda int, s, scond, amax []float64, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zheequb_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cheev.f.
func Cheev(jobz lapack.Job, ul blas.Uplo, n int, a []complex64, lda int, w []float32, work []complex64, lwork int, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cheev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_w), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zheev.f.
func Zheev(jobz lapack.Job, ul blas.Uplo, n int, a []complex128, lda int, w []float64, work []complex128, lwork int, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zheev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_w), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cheevd.f.
func Cheevd(jobz lapack.Job, ul blas.Uplo, n int, a []complex64, lda int, w []float32, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cheevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_w), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zheevd.f.
func Zheevd(jobz lapack.Job, ul blas.Uplo, n int, a []complex128, lda int, w []float64, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zheevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_w), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cheevr.f.
func Cheevr(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []complex64, lda int, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, isuppz []int32, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cheevr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zheevr.f.
func Zheevr(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []complex128, lda int, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, isuppz []int32, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zheevr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cheevx.f.
func Cheevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []complex64, lda int, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_cheevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zheevx.f.
func Zheevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []complex128, lda int, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_zheevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chegst.f.
func Chegst(itype int, ul blas.Uplo, n int, a []complex64, lda int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_chegst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhegst.f.
func Zhegst(itype int, ul blas.Uplo, n int, a []complex128, lda int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zhegst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chegv.f.
func Chegv(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []complex64, lda int, b []complex64, ldb int, w []float32, work []complex64, lwork int, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chegv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.float)(_w), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhegv.f.
func Zhegv(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []complex128, lda int, b []complex128, ldb int, w []float64, work []complex128, lwork int, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhegv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.double)(_w), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chegvd.f.
func Chegvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []complex64, lda int, b []complex64, ldb int, w []float32, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_chegvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.float)(_w), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhegvd.f.
func Zhegvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []complex128, lda int, b []complex128, ldb int, w []float64, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zhegvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.double)(_w), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chegvx.f.
func Chegvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []complex64, lda int, b []complex64, ldb int, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_chegvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhegvx.f.
func Zhegvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []complex128, lda int, b []complex128, ldb int, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_zhegvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cherfs.f.
func Cherfs(ul blas.Uplo, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cherfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zherfs.f.
func Zherfs(ul blas.Uplo, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zherfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chesv.f.
func Chesv(ul blas.Uplo, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chesv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhesv.f.
func Zhesv(ul blas.Uplo, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhesv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chesvx.f.
func Chesvx(fact byte, ul blas.Uplo, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, lwork int, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chesvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhesvx.f.
func Zhesvx(fact byte, ul blas.Uplo, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, lwork int, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhesvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetrd.f.
func Chetrd(ul blas.Uplo, n int, a []complex64, lda int, d, e []float32, tau, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chetrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetrd.f.
func Zhetrd(ul blas.Uplo, n int, a []complex128, lda int, d, e []float64, tau, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhetrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetrf.f.
func Chetrf(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chetrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetrf.f.
func Zhetrf(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhetrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetri.f.
func Chetri(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chetri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetri.f.
func Zhetri(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhetri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetrs.f.
func Chetrs(ul blas.Uplo, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_chetrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetrs.f.
func Zhetrs(ul blas.Uplo, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zhetrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chfrk.f.
func Chfrk(transr blas.Transpose, ul blas.Uplo, trans blas.Transpose, n, k int, alpha float32, a []complex64, lda int, beta float32, c []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	return isZero(C.LAPACKE_chfrk_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(k), (C.float)(alpha), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.float)(beta), (*C.lapack_complex_float)(_c)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhfrk.f.
func Zhfrk(transr blas.Transpose, ul blas.Uplo, trans blas.Transpose, n, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	return isZero(C.LAPACKE_zhfrk_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(k), (C.double)(alpha), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.double)(beta), (*C.lapack_complex_double)(_c)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/shgeqz.f.
func Shgeqz(job lapack.Job, compq, compz lapack.Comp, n, ilo, ihi int, h []float32, ldh int, t []float32, ldt int, alphar, alphai, beta, q []float32, ldq int, z []float32, ldz int, work []float32, lwork int) bool {
	var _h *float32
	if len(h) > 0 {
		_h = &h[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _alphar *float32
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float32
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_shgeqz_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_h), (C.lapack_int)(ldh), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_alphar), (*C.float)(_alphai), (*C.float)(_beta), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dhgeqz.f.
func Dhgeqz(job lapack.Job, compq, compz lapack.Comp, n, ilo, ihi int, h []float64, ldh int, t []float64, ldt int, alphar, alphai, beta, q []float64, ldq int, z []float64, ldz int, work []float64, lwork int) bool {
	var _h *float64
	if len(h) > 0 {
		_h = &h[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _alphar *float64
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float64
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dhgeqz_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_h), (C.lapack_int)(ldh), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_alphar), (*C.double)(_alphai), (*C.double)(_beta), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chgeqz.f.
func Chgeqz(job lapack.Job, compq, compz lapack.Comp, n, ilo, ihi int, h []complex64, ldh int, t []complex64, ldt int, alpha, beta, q []complex64, ldq int, z []complex64, ldz int, work []complex64, lwork int, rwork []float32) bool {
	var _h *complex64
	if len(h) > 0 {
		_h = &h[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _alpha *complex64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chgeqz_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_h), (C.lapack_int)(ldh), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_alpha), (*C.lapack_complex_float)(_beta), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhgeqz.f.
func Zhgeqz(job lapack.Job, compq, compz lapack.Comp, n, ilo, ihi int, h []complex128, ldh int, t []complex128, ldt int, alpha, beta, q []complex128, ldq int, z []complex128, ldz int, work []complex128, lwork int, rwork []float64) bool {
	var _h *complex128
	if len(h) > 0 {
		_h = &h[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _alpha *complex128
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex128
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhgeqz_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_h), (C.lapack_int)(ldh), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_alpha), (*C.lapack_complex_double)(_beta), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpcon.f.
func Chpcon(ul blas.Uplo, n int, ap []complex64, ipiv []int32, anorm float32, rcond []float32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chpcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpcon.f.
func Zhpcon(ul blas.Uplo, n int, ap []complex128, ipiv []int32, anorm float64, rcond []float64, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhpcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpev.f.
func Chpev(jobz lapack.Job, ul blas.Uplo, n int, ap []complex64, w []float32, z []complex64, ldz int, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chpev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpev.f.
func Zhpev(jobz lapack.Job, ul blas.Uplo, n int, ap []complex128, w []float64, z []complex128, ldz int, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhpev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpevd.f.
func Chpevd(jobz lapack.Job, ul blas.Uplo, n int, ap []complex64, w []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_chpevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpevd.f.
func Zhpevd(jobz lapack.Job, ul blas.Uplo, n int, ap []complex128, w []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zhpevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpevx.f.
func Chpevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap []complex64, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, work []complex64, rwork []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_chpevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpevx.f.
func Zhpevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap []complex128, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, work []complex128, rwork []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_zhpevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpgst.f.
func Chpgst(itype int, ul blas.Uplo, n int, ap, bp []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	return isZero(C.LAPACKE_chpgst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_bp)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpgst.f.
func Zhpgst(itype int, ul blas.Uplo, n int, ap, bp []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex128
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	return isZero(C.LAPACKE_zhpgst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_bp)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpgv.f.
func Chpgv(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp []complex64, w []float32, z []complex64, ldz int, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chpgv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_bp), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpgv.f.
func Zhpgv(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp []complex128, w []float64, z []complex128, ldz int, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex128
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhpgv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_bp), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpgvd.f.
func Chpgvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp []complex64, w []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_chpgvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_bp), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpgvd.f.
func Zhpgvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp []complex128, w []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex128
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zhpgvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_bp), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpgvx.f.
func Chpgvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap, bp []complex64, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, work []complex64, rwork []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_chpgvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_bp), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (*C.float)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpgvx.f.
func Zhpgvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap, bp []complex128, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, work []complex128, rwork []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *complex128
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_zhpgvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_bp), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (*C.double)(_rwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chprfs.f.
func Chprfs(ul blas.Uplo, n, nrhs int, ap, afp []complex64, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhprfs.f.
func Zhprfs(ul blas.Uplo, n, nrhs int, ap, afp []complex128, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex128
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpsv.f.
func Chpsv(ul blas.Uplo, n, nrhs int, ap []complex64, ipiv []int32, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_chpsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpsv.f.
func Zhpsv(ul blas.Uplo, n, nrhs int, ap []complex128, ipiv []int32, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zhpsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chpsvx.f.
func Chpsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []complex64, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_chpsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhpsvx.f.
func Zhpsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []complex128, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex128
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zhpsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chptrd.f.
func Chptrd(ul blas.Uplo, n int, ap []complex64, d, e []float32, tau []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_chptrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhptrd.f.
func Zhptrd(ul blas.Uplo, n int, ap []complex128, d, e []float64, tau []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_zhptrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chptrf.f.
func Chptrf(ul blas.Uplo, n int, ap []complex64, ipiv []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_chptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhptrf.f.
func Zhptrf(ul blas.Uplo, n int, ap []complex128, ipiv []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zhptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chptri.f.
func Chptri(ul blas.Uplo, n int, ap []complex64, ipiv []int32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhptri.f.
func Zhptri(ul blas.Uplo, n int, ap []complex128, ipiv []int32, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chptrs.f.
func Chptrs(ul blas.Uplo, n, nrhs int, ap []complex64, ipiv []int32, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_chptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhptrs.f.
func Zhptrs(ul blas.Uplo, n, nrhs int, ap []complex128, ipiv []int32, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zhptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/shsein.f.
func Shsein(job lapack.Job, eigsrc, initv byte, sel []int32, n int, h []float32, ldh int, wr, wi, vl []float32, ldvl int, vr []float32, ldvr, mm int, m []int32, work []float32, ifaill, ifailr []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _h *float32
	if len(h) > 0 {
		_h = &h[0]
	}
	var _wr *float32
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float32
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _ifaill *int32
	if len(ifaill) > 0 {
		_ifaill = &ifaill[0]
	}
	var _ifailr *int32
	if len(ifailr) > 0 {
		_ifailr = &ifailr[0]
	}
	return isZero(C.LAPACKE_shsein_work((C.int)(rowMajor), (C.char)(job), (C.char)(eigsrc), (C.char)(initv), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_h), (C.lapack_int)(ldh), (*C.float)(_wr), (*C.float)(_wi), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.float)(_work), (*C.lapack_int)(_ifaill), (*C.lapack_int)(_ifailr)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dhsein.f.
func Dhsein(job lapack.Job, eigsrc, initv byte, sel []int32, n int, h []float64, ldh int, wr, wi, vl []float64, ldvl int, vr []float64, ldvr, mm int, m []int32, work []float64, ifaill, ifailr []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _h *float64
	if len(h) > 0 {
		_h = &h[0]
	}
	var _wr *float64
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float64
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _ifaill *int32
	if len(ifaill) > 0 {
		_ifaill = &ifaill[0]
	}
	var _ifailr *int32
	if len(ifailr) > 0 {
		_ifailr = &ifailr[0]
	}
	return isZero(C.LAPACKE_dhsein_work((C.int)(rowMajor), (C.char)(job), (C.char)(eigsrc), (C.char)(initv), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_h), (C.lapack_int)(ldh), (*C.double)(_wr), (*C.double)(_wi), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.double)(_work), (*C.lapack_int)(_ifaill), (*C.lapack_int)(_ifailr)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chsein.f.
func Chsein(job lapack.Job, eigsrc, initv byte, sel []int32, n int, h []complex64, ldh int, w, vl []complex64, ldvl int, vr []complex64, ldvr, mm int, m []int32, work []complex64, rwork []float32, ifaill, ifailr []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _h *complex64
	if len(h) > 0 {
		_h = &h[0]
	}
	var _w *complex64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _ifaill *int32
	if len(ifaill) > 0 {
		_ifaill = &ifaill[0]
	}
	var _ifailr *int32
	if len(ifailr) > 0 {
		_ifailr = &ifailr[0]
	}
	return isZero(C.LAPACKE_chsein_work((C.int)(rowMajor), (C.char)(job), (C.char)(eigsrc), (C.char)(initv), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_h), (C.lapack_int)(ldh), (*C.lapack_complex_float)(_w), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_float)(_work), (*C.float)(_rwork), (*C.lapack_int)(_ifaill), (*C.lapack_int)(_ifailr)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhsein.f.
func Zhsein(job lapack.Job, eigsrc, initv byte, sel []int32, n int, h []complex128, ldh int, w, vl []complex128, ldvl int, vr []complex128, ldvr, mm int, m []int32, work []complex128, rwork []float64, ifaill, ifailr []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _h *complex128
	if len(h) > 0 {
		_h = &h[0]
	}
	var _w *complex128
	if len(w) > 0 {
		_w = &w[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _ifaill *int32
	if len(ifaill) > 0 {
		_ifaill = &ifaill[0]
	}
	var _ifailr *int32
	if len(ifailr) > 0 {
		_ifailr = &ifailr[0]
	}
	return isZero(C.LAPACKE_zhsein_work((C.int)(rowMajor), (C.char)(job), (C.char)(eigsrc), (C.char)(initv), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_h), (C.lapack_int)(ldh), (*C.lapack_complex_double)(_w), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_double)(_work), (*C.double)(_rwork), (*C.lapack_int)(_ifaill), (*C.lapack_int)(_ifailr)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/shseqr.f.
func Shseqr(job lapack.Job, compz lapack.Comp, n, ilo, ihi int, h []float32, ldh int, wr, wi, z []float32, ldz int, work []float32, lwork int) int {
	var _h *float32
	if len(h) > 0 {
		_h = &h[0]
	}
	var _wr *float32
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float32
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return int(C.LAPACKE_shseqr_work((C.int)(rowMajor), (C.char)(job), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_h), (C.lapack_int)(ldh), (*C.float)(_wr), (*C.float)(_wi), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dhseqr.f.
func Dhseqr(job lapack.Job, compz lapack.Comp, n, ilo, ihi int, h []float64, ldh int, wr, wi, z []float64, ldz int, work []float64, lwork int) int {
	var _h *float64
	if len(h) > 0 {
		_h = &h[0]
	}
	var _wr *float64
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float64
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return int(C.LAPACKE_dhseqr_work((C.int)(rowMajor), (C.char)(job), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_h), (C.lapack_int)(ldh), (*C.double)(_wr), (*C.double)(_wi), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chseqr.f.
func Chseqr(job lapack.Job, compz lapack.Comp, n, ilo, ihi int, h []complex64, ldh int, w, z []complex64, ldz int, work []complex64, lwork int) int {
	var _h *complex64
	if len(h) > 0 {
		_h = &h[0]
	}
	var _w *complex64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return int(C.LAPACKE_chseqr_work((C.int)(rowMajor), (C.char)(job), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_h), (C.lapack_int)(ldh), (*C.lapack_complex_float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhseqr.f.
func Zhseqr(job lapack.Job, compz lapack.Comp, n, ilo, ihi int, h []complex128, ldh int, w, z []complex128, ldz int, work []complex128, lwork int) int {
	var _h *complex128
	if len(h) > 0 {
		_h = &h[0]
	}
	var _w *complex128
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return int(C.LAPACKE_zhseqr_work((C.int)(rowMajor), (C.char)(job), (C.char)(compz), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_h), (C.lapack_int)(ldh), (*C.lapack_complex_double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clacgv.f.
func Clacgv(n int, x []complex64, incx int) bool {
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	return isZero(C.LAPACKE_clacgv_work((C.lapack_int)(n), (*C.lapack_complex_float)(_x), (C.lapack_int)(incx)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlacgv.f.
func Zlacgv(n int, x []complex128, incx int) bool {
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	return isZero(C.LAPACKE_zlacgv_work((C.lapack_int)(n), (*C.lapack_complex_double)(_x), (C.lapack_int)(incx)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slacn2.f.
func Slacn2(n int, v, x []float32, isgn []int32, est []float32, kase, isave []int32) bool {
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _isgn *int32
	if len(isgn) > 0 {
		_isgn = &isgn[0]
	}
	var _est *float32
	if len(est) > 0 {
		_est = &est[0]
	}
	var _kase *int32
	if len(kase) > 0 {
		_kase = &kase[0]
	}
	var _isave *int32
	if len(isave) > 0 {
		_isave = &isave[0]
	}
	return isZero(C.LAPACKE_slacn2_work((C.lapack_int)(n), (*C.float)(_v), (*C.float)(_x), (*C.lapack_int)(_isgn), (*C.float)(_est), (*C.lapack_int)(_kase), (*C.lapack_int)(_isave)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlacn2.f.
func Dlacn2(n int, v, x []float64, isgn []int32, est []float64, kase, isave []int32) bool {
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _isgn *int32
	if len(isgn) > 0 {
		_isgn = &isgn[0]
	}
	var _est *float64
	if len(est) > 0 {
		_est = &est[0]
	}
	var _kase *int32
	if len(kase) > 0 {
		_kase = &kase[0]
	}
	var _isave *int32
	if len(isave) > 0 {
		_isave = &isave[0]
	}
	return isZero(C.LAPACKE_dlacn2_work((C.lapack_int)(n), (*C.double)(_v), (*C.double)(_x), (*C.lapack_int)(_isgn), (*C.double)(_est), (*C.lapack_int)(_kase), (*C.lapack_int)(_isave)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clacn2.f.
func Clacn2(n int, v, x []complex64, est []float32, kase, isave []int32) bool {
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _est *float32
	if len(est) > 0 {
		_est = &est[0]
	}
	var _kase *int32
	if len(kase) > 0 {
		_kase = &kase[0]
	}
	var _isave *int32
	if len(isave) > 0 {
		_isave = &isave[0]
	}
	return isZero(C.LAPACKE_clacn2_work((C.lapack_int)(n), (*C.lapack_complex_float)(_v), (*C.lapack_complex_float)(_x), (*C.float)(_est), (*C.lapack_int)(_kase), (*C.lapack_int)(_isave)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlacn2.f.
func Zlacn2(n int, v, x []complex128, est []float64, kase, isave []int32) bool {
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _est *float64
	if len(est) > 0 {
		_est = &est[0]
	}
	var _kase *int32
	if len(kase) > 0 {
		_kase = &kase[0]
	}
	var _isave *int32
	if len(isave) > 0 {
		_isave = &isave[0]
	}
	return isZero(C.LAPACKE_zlacn2_work((C.lapack_int)(n), (*C.lapack_complex_double)(_v), (*C.lapack_complex_double)(_x), (*C.double)(_est), (*C.lapack_int)(_kase), (*C.lapack_int)(_isave)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slacpy.f.
func Slacpy(ul blas.Uplo, m, n int, a []float32, lda int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	case blas.All:
		ul = 'A'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_slacpy_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlacpy.f.
func Dlacpy(ul blas.Uplo, m, n int, a []float64, lda int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	case blas.All:
		ul = 'A'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dlacpy_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clacpy.f.
func Clacpy(ul blas.Uplo, m, n int, a []complex64, lda int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	case blas.All:
		ul = 'A'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_clacpy_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlacpy.f.
func Zlacpy(ul blas.Uplo, m, n int, a []complex128, lda int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	case blas.All:
		ul = 'A'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zlacpy_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clacp2.f.
func Clacp2(ul blas.Uplo, m, n int, a []float32, lda int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_clacp2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlacp2.f.
func Zlacp2(ul blas.Uplo, m, n int, a []float64, lda int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zlacp2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlag2c.f.
func Zlag2c(m, n int, a []complex128, lda int, sa []complex64, ldsa int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sa *complex64
	if len(sa) > 0 {
		_sa = &sa[0]
	}
	return isZero(C.LAPACKE_zlag2c_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_sa), (C.lapack_int)(ldsa)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slag2d.f.
func Slag2d(m, n int, sa []float32, ldsa int, a []float64, lda int) bool {
	var _sa *float32
	if len(sa) > 0 {
		_sa = &sa[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_slag2d_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_sa), (C.lapack_int)(ldsa), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlag2s.f.
func Dlag2s(m, n int, a []float64, lda int, sa []float32, ldsa int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _sa *float32
	if len(sa) > 0 {
		_sa = &sa[0]
	}
	return isZero(C.LAPACKE_dlag2s_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.float)(_sa), (C.lapack_int)(ldsa)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clag2z.f.
func Clag2z(m, n int, sa []complex64, ldsa int, a []complex128, lda int) bool {
	var _sa *complex64
	if len(sa) > 0 {
		_sa = &sa[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_clag2z_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_sa), (C.lapack_int)(ldsa), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slagge.f.
func Slagge(m, n, kl, ku int, d, a []float32, lda int, iseed []int32, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_slagge_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_d), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlagge.f.
func Dlagge(m, n, kl, ku int, d, a []float64, lda int, iseed []int32, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dlagge_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_d), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clagge.f.
func Clagge(m, n, kl, ku int, d []float32, a []complex64, lda int, iseed []int32, work []complex64) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_clagge_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.float)(_d), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlagge.f.
func Zlagge(m, n, kl, ku int, d []float64, a []complex128, lda int, iseed []int32, work []complex128) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zlagge_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(kl), (C.lapack_int)(ku), (*C.double)(_d), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/claghe.f.
func Claghe(n, k int, d []float32, a []complex64, lda int, iseed []int32, work []complex64) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_claghe_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_d), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlaghe.f.
func Zlaghe(n, k int, d []float64, a []complex128, lda int, iseed []int32, work []complex128) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zlaghe_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_d), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slagsy.f.
func Slagsy(n, k int, d, a []float32, lda int, iseed []int32, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_slagsy_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_d), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlagsy.f.
func Dlagsy(n, k int, d, a []float64, lda int, iseed []int32, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dlagsy_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_d), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clagsy.f.
func Clagsy(n, k int, d []float32, a []complex64, lda int, iseed []int32, work []complex64) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_clagsy_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_d), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlagsy.f.
func Zlagsy(n, k int, d []float64, a []complex128, lda int, iseed []int32, work []complex128) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zlagsy_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_d), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_iseed), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slapmr.f.
func Slapmr(forwrd int32, m, n int, x []float32, ldx int, k []int32) bool {
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_slapmr_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlapmr.f.
func Dlapmr(forwrd int32, m, n int, x []float64, ldx int, k []int32) bool {
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_dlapmr_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clapmr.f.
func Clapmr(forwrd int32, m, n int, x []complex64, ldx int, k []int32) bool {
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_clapmr_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlapmr.f.
func Zlapmr(forwrd int32, m, n int, x []complex128, ldx int, k []int32) bool {
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_zlapmr_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slapmt.f.
func Slapmt(forwrd int32, m, n int, x []float32, ldx int, k []int32) bool {
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_slapmt_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlapmt.f.
func Dlapmt(forwrd int32, m, n int, x []float64, ldx int, k []int32) bool {
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_dlapmt_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clapmt.f.
func Clapmt(forwrd int32, m, n int, x []complex64, ldx int, k []int32) bool {
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_clapmt_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlapmt.f.
func Zlapmt(forwrd int32, m, n int, x []complex128, ldx int, k []int32) bool {
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _k *int32
	if len(k) > 0 {
		_k = &k[0]
	}
	return isZero(C.LAPACKE_zlapmt_work((C.int)(rowMajor), (C.lapack_logical)(forwrd), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.lapack_int)(_k)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slartgp.f.
func Slartgp(f, g float32, cs, sn, r []float32) bool {
	var _cs *float32
	if len(cs) > 0 {
		_cs = &cs[0]
	}
	var _sn *float32
	if len(sn) > 0 {
		_sn = &sn[0]
	}
	var _r *float32
	if len(r) > 0 {
		_r = &r[0]
	}
	return isZero(C.LAPACKE_slartgp_work((C.float)(f), (C.float)(g), (*C.float)(_cs), (*C.float)(_sn), (*C.float)(_r)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlartgp.f.
func Dlartgp(f, g float64, cs, sn, r []float64) bool {
	var _cs *float64
	if len(cs) > 0 {
		_cs = &cs[0]
	}
	var _sn *float64
	if len(sn) > 0 {
		_sn = &sn[0]
	}
	var _r *float64
	if len(r) > 0 {
		_r = &r[0]
	}
	return isZero(C.LAPACKE_dlartgp_work((C.double)(f), (C.double)(g), (*C.double)(_cs), (*C.double)(_sn), (*C.double)(_r)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slartgs.f.
func Slartgs(x, y, sigma float32, cs, sn []float32) bool {
	var _cs *float32
	if len(cs) > 0 {
		_cs = &cs[0]
	}
	var _sn *float32
	if len(sn) > 0 {
		_sn = &sn[0]
	}
	return isZero(C.LAPACKE_slartgs_work((C.float)(x), (C.float)(y), (C.float)(sigma), (*C.float)(_cs), (*C.float)(_sn)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlartgs.f.
func Dlartgs(x, y, sigma float64, cs, sn []float64) bool {
	var _cs *float64
	if len(cs) > 0 {
		_cs = &cs[0]
	}
	var _sn *float64
	if len(sn) > 0 {
		_sn = &sn[0]
	}
	return isZero(C.LAPACKE_dlartgs_work((C.double)(x), (C.double)(y), (C.double)(sigma), (*C.double)(_cs), (*C.double)(_sn)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slapy2.f.
func Slapy2(x, y float32) float32 {
	return float32(C.LAPACKE_slapy2_work((C.float)(x), (C.float)(y)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlapy2.f.
func Dlapy2(x, y float64) float64 {
	return float64(C.LAPACKE_dlapy2_work((C.double)(x), (C.double)(y)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slapy3.f.
func Slapy3(x, y, z float32) float32 {
	return float32(C.LAPACKE_slapy3_work((C.float)(x), (C.float)(y), (C.float)(z)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlapy3.f.
func Dlapy3(x, y, z float64) float64 {
	return float64(C.LAPACKE_dlapy3_work((C.double)(x), (C.double)(y), (C.double)(z)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slamch.f.
func Slamch(cmach byte) float32 {
	return float32(C.LAPACKE_slamch_work((C.char)(cmach)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlamch.f.
func Dlamch(cmach byte) float64 {
	return float64(C.LAPACKE_dlamch_work((C.char)(cmach)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slange.f.
func Slange(norm byte, m, n int, a []float32, lda int, work []float32) float32 {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_slange_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlange.f.
func Dlange(norm byte, m, n int, a []float64, lda int, work []float64) float64 {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_dlange_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clange.f.
func Clange(norm byte, m, n int, a []complex64, lda int, work []float32) float32 {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_clange_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlange.f.
func Zlange(norm byte, m, n int, a []complex128, lda int, work []float64) float64 {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_zlange_work((C.int)(rowMajor), (C.char)(norm), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clanhe.f.
func Clanhe(norm byte, ul blas.Uplo, n int, a []complex64, lda int, work []float32) float32 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_clanhe_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlanhe.f.
func Zlanhe(norm byte, ul blas.Uplo, n int, a []complex128, lda int, work []float64) float64 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_zlanhe_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slansy.f.
func Slansy(norm byte, ul blas.Uplo, n int, a []float32, lda int, work []float32) float32 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_slansy_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlansy.f.
func Dlansy(norm byte, ul blas.Uplo, n int, a []float64, lda int, work []float64) float64 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_dlansy_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clansy.f.
func Clansy(norm byte, ul blas.Uplo, n int, a []complex64, lda int, work []float32) float32 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_clansy_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlansy.f.
func Zlansy(norm byte, ul blas.Uplo, n int, a []complex128, lda int, work []float64) float64 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_zlansy_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slantr.f.
func Slantr(norm byte, ul blas.Uplo, d blas.Diag, m, n int, a []float32, lda int, work []float32) float32 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_slantr_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlantr.f.
func Dlantr(norm byte, ul blas.Uplo, d blas.Diag, m, n int, a []float64, lda int, work []float64) float64 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_dlantr_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clantr.f.
func Clantr(norm byte, ul blas.Uplo, d blas.Diag, m, n int, a []complex64, lda int, work []float32) float32 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return float32(C.LAPACKE_clantr_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlantr.f.
func Zlantr(norm byte, ul blas.Uplo, d blas.Diag, m, n int, a []complex128, lda int, work []float64) float64 {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return float64(C.LAPACKE_zlantr_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slarfb.f.
func Slarfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k int, v []float32, ldv int, t []float32, ldt int, c []float32, ldc int, work []float32, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_slarfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlarfb.f.
func Dlarfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k int, v []float64, ldv int, t []float64, ldt int, c []float64, ldc int, work []float64, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dlarfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clarfb.f.
func Clarfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k int, v []complex64, ldv int, t []complex64, ldt int, c []complex64, ldc int, work []complex64, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_clarfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlarfb.f.
func Zlarfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k int, v []complex128, ldv int, t []complex128, ldt int, c []complex128, ldc int, work []complex128, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zlarfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slarfg.f.
func Slarfg(n int, alpha, x []float32, incx int, tau []float32) bool {
	var _alpha *float32
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_slarfg_work((C.lapack_int)(n), (*C.float)(_alpha), (*C.float)(_x), (C.lapack_int)(incx), (*C.float)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlarfg.f.
func Dlarfg(n int, alpha, x []float64, incx int, tau []float64) bool {
	var _alpha *float64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_dlarfg_work((C.lapack_int)(n), (*C.double)(_alpha), (*C.double)(_x), (C.lapack_int)(incx), (*C.double)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clarfg.f.
func Clarfg(n int, alpha, x []complex64, incx int, tau []complex64) bool {
	var _alpha *complex64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_clarfg_work((C.lapack_int)(n), (*C.lapack_complex_float)(_alpha), (*C.lapack_complex_float)(_x), (C.lapack_int)(incx), (*C.lapack_complex_float)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlarfg.f.
func Zlarfg(n int, alpha, x []complex128, incx int, tau []complex128) bool {
	var _alpha *complex128
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_zlarfg_work((C.lapack_int)(n), (*C.lapack_complex_double)(_alpha), (*C.lapack_complex_double)(_x), (C.lapack_int)(incx), (*C.lapack_complex_double)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slarft.f.
func Slarft(direct, storev byte, n, k int, v []float32, ldv int, tau, t []float32, ldt int) bool {
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_slarft_work((C.int)(rowMajor), (C.char)(direct), (C.char)(storev), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_tau), (*C.float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlarft.f.
func Dlarft(direct, storev byte, n, k int, v []float64, ldv int, tau, t []float64, ldt int) bool {
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_dlarft_work((C.int)(rowMajor), (C.char)(direct), (C.char)(storev), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_tau), (*C.double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clarft.f.
func Clarft(direct, storev byte, n, k int, v []complex64, ldv int, tau, t []complex64, ldt int) bool {
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_clarft_work((C.int)(rowMajor), (C.char)(direct), (C.char)(storev), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlarft.f.
func Zlarft(direct, storev byte, n, k int, v []complex128, ldv int, tau, t []complex128, ldt int) bool {
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_zlarft_work((C.int)(rowMajor), (C.char)(direct), (C.char)(storev), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slarfx.f.
func Slarfx(side blas.Side, m, n int, v []float32, tau float32, c []float32, ldc int, work []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_slarfx_work((C.int)(rowMajor), (C.char)(side), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_v), (C.float)(tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlarfx.f.
func Dlarfx(side blas.Side, m, n int, v []float64, tau float64, c []float64, ldc int, work []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dlarfx_work((C.int)(rowMajor), (C.char)(side), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_v), (C.double)(tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clarfx.f.
func Clarfx(side blas.Side, m, n int, v []complex64, tau complex64, c []complex64, ldc int, work []complex64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_clarfx_work((C.int)(rowMajor), (C.char)(side), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_v), (C.lapack_complex_float)(tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlarfx.f.
func Zlarfx(side blas.Side, m, n int, v []complex128, tau complex128, c []complex128, ldc int, work []complex128) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zlarfx_work((C.int)(rowMajor), (C.char)(side), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_v), (C.lapack_complex_double)(tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slarnv.f.
func Slarnv(idist int, iseed []int32, n int, x []float32) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	return isZero(C.LAPACKE_slarnv_work((C.lapack_int)(idist), (*C.lapack_int)(_iseed), (C.lapack_int)(n), (*C.float)(_x)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlarnv.f.
func Dlarnv(idist int, iseed []int32, n int, x []float64) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	return isZero(C.LAPACKE_dlarnv_work((C.lapack_int)(idist), (*C.lapack_int)(_iseed), (C.lapack_int)(n), (*C.double)(_x)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clarnv.f.
func Clarnv(idist int, iseed []int32, n int, x []complex64) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	return isZero(C.LAPACKE_clarnv_work((C.lapack_int)(idist), (*C.lapack_int)(_iseed), (C.lapack_int)(n), (*C.lapack_complex_float)(_x)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlarnv.f.
func Zlarnv(idist int, iseed []int32, n int, x []complex128) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	return isZero(C.LAPACKE_zlarnv_work((C.lapack_int)(idist), (*C.lapack_int)(_iseed), (C.lapack_int)(n), (*C.lapack_complex_double)(_x)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slascl.f.
func Slascl(typ byte, kl, ku int, cfrom, cto float32, m, n int, a []float32, lda int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_slascl_work((C.int)(rowMajor), (C.char)(typ), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.float)(cfrom), (C.float)(cto), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlascl.f.
func Dlascl(typ byte, kl, ku int, cfrom, cto float64, m, n int, a []float64, lda int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dlascl_work((C.int)(rowMajor), (C.char)(typ), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.double)(cfrom), (C.double)(cto), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clascl.f.
func Clascl(typ byte, kl, ku int, cfrom, cto float32, m, n int, a []complex64, lda int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_clascl_work((C.int)(rowMajor), (C.char)(typ), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.float)(cfrom), (C.float)(cto), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlascl.f.
func Zlascl(typ byte, kl, ku int, cfrom, cto float64, m, n int, a []complex128, lda int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zlascl_work((C.int)(rowMajor), (C.char)(typ), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.double)(cfrom), (C.double)(cto), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slaset.f.
func Slaset(ul blas.Uplo, m, n int, alpha, beta float32, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_slaset_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (C.float)(alpha), (C.float)(beta), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlaset.f.
func Dlaset(ul blas.Uplo, m, n int, alpha, beta float64, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dlaset_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (C.double)(alpha), (C.double)(beta), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/claset.f.
func Claset(ul blas.Uplo, m, n int, alpha, beta complex64, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_claset_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_complex_float)(alpha), (C.lapack_complex_float)(beta), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlaset.f.
func Zlaset(ul blas.Uplo, m, n int, alpha, beta complex128, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zlaset_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_complex_double)(alpha), (C.lapack_complex_double)(beta), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slasrt.f.
func Slasrt(id byte, n int, d []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	return isZero(C.LAPACKE_slasrt_work((C.char)(id), (C.lapack_int)(n), (*C.float)(_d)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlasrt.f.
func Dlasrt(id byte, n int, d []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	return isZero(C.LAPACKE_dlasrt_work((C.char)(id), (C.lapack_int)(n), (*C.double)(_d)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slaswp.f.
func Slaswp(n int, a []float32, lda, k1, k2 int, ipiv []int32, incx int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_slaswp_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (C.lapack_int)(k1), (C.lapack_int)(k2), (*C.lapack_int)(_ipiv), (C.lapack_int)(incx)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlaswp.f.
func Dlaswp(n int, a []float64, lda, k1, k2 int, ipiv []int32, incx int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dlaswp_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (C.lapack_int)(k1), (C.lapack_int)(k2), (*C.lapack_int)(_ipiv), (C.lapack_int)(incx)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/claswp.f.
func Claswp(n int, a []complex64, lda, k1, k2 int, ipiv []int32, incx int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_claswp_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.lapack_int)(k1), (C.lapack_int)(k2), (*C.lapack_int)(_ipiv), (C.lapack_int)(incx)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlaswp.f.
func Zlaswp(n int, a []complex128, lda, k1, k2 int, ipiv []int32, incx int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zlaswp_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.lapack_int)(k1), (C.lapack_int)(k2), (*C.lapack_int)(_ipiv), (C.lapack_int)(incx)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slatms.f.
func Slatms(m, n int, dist byte, iseed []int32, sym byte, d []float32, mode int, cond, dmax float32, kl, ku int, pack byte, a []float32, lda int, work []float32) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_slatms_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.char)(dist), (*C.lapack_int)(_iseed), (C.char)(sym), (*C.float)(_d), (C.lapack_int)(mode), (C.float)(cond), (C.float)(dmax), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.char)(pack), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlatms.f.
func Dlatms(m, n int, dist byte, iseed []int32, sym byte, d []float64, mode int, cond, dmax float64, kl, ku int, pack byte, a []float64, lda int, work []float64) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dlatms_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.char)(dist), (*C.lapack_int)(_iseed), (C.char)(sym), (*C.double)(_d), (C.lapack_int)(mode), (C.double)(cond), (C.double)(dmax), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.char)(pack), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clatms.f.
func Clatms(m, n int, dist byte, iseed []int32, sym byte, d []float32, mode int, cond, dmax float32, kl, ku int, pack byte, a []complex64, lda int, work []complex64) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_clatms_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.char)(dist), (*C.lapack_int)(_iseed), (C.char)(sym), (*C.float)(_d), (C.lapack_int)(mode), (C.float)(cond), (C.float)(dmax), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.char)(pack), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlatms.f.
func Zlatms(m, n int, dist byte, iseed []int32, sym byte, d []float64, mode int, cond, dmax float64, kl, ku int, pack byte, a []complex128, lda int, work []complex128) bool {
	var _iseed *int32
	if len(iseed) > 0 {
		_iseed = &iseed[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zlatms_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.char)(dist), (*C.lapack_int)(_iseed), (C.char)(sym), (*C.double)(_d), (C.lapack_int)(mode), (C.double)(cond), (C.double)(dmax), (C.lapack_int)(kl), (C.lapack_int)(ku), (C.char)(pack), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/slauum.f.
func Slauum(ul blas.Uplo, n int, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_slauum_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dlauum.f.
func Dlauum(ul blas.Uplo, n int, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dlauum_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/clauum.f.
func Clauum(ul blas.Uplo, n int, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_clauum_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zlauum.f.
func Zlauum(ul blas.Uplo, n int, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zlauum_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sopgtr.f.
func Sopgtr(ul blas.Uplo, n int, ap, tau, q []float32, ldq int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sopgtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_tau), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dopgtr.f.
func Dopgtr(ul blas.Uplo, n int, ap, tau, q []float64, ldq int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dopgtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_tau), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sopmtr.f.
func Sopmtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, ap, tau, c []float32, ldc int, work []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sopmtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dopmtr.f.
func Dopmtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, ap, tau, c []float64, ldc int, work []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dopmtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorgbr.f.
func Sorgbr(vect byte, m, n, k int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorgbr_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorgbr.f.
func Dorgbr(vect byte, m, n, k int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorgbr_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorghr.f.
func Sorghr(n, ilo, ihi int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorghr_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorghr.f.
func Dorghr(n, ilo, ihi int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorghr_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorglq.f.
func Sorglq(m, n, k int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorglq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorglq.f.
func Dorglq(m, n, k int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorglq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorgql.f.
func Sorgql(m, n, k int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorgql_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorgql.f.
func Dorgql(m, n, k int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorgql_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorgqr.f.
func Sorgqr(m, n, k int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorgqr_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorgqr.f.
func Dorgqr(m, n, k int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorgqr_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorgrq.f.
func Sorgrq(m, n, k int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorgrq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorgrq.f.
func Dorgrq(m, n, k int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorgrq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorgtr.f.
func Sorgtr(ul blas.Uplo, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorgtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorgtr.f.
func Dorgtr(ul blas.Uplo, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorgtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormbr.f.
func Sormbr(vect byte, side blas.Side, trans blas.Transpose, m, n, k int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormbr_work((C.int)(rowMajor), (C.char)(vect), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormbr.f.
func Dormbr(vect byte, side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormbr_work((C.int)(rowMajor), (C.char)(vect), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormhr.f.
func Sormhr(side blas.Side, trans blas.Transpose, m, n, ilo, ihi int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormhr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormhr.f.
func Dormhr(side blas.Side, trans blas.Transpose, m, n, ilo, ihi int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormhr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormlq.f.
func Sormlq(side blas.Side, trans blas.Transpose, m, n, k int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormlq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormlq.f.
func Dormlq(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormlq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormql.f.
func Sormql(side blas.Side, trans blas.Transpose, m, n, k int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormql_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormql.f.
func Dormql(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormql_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormqr.f.
func Sormqr(side blas.Side, trans blas.Transpose, m, n, k int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormqr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormqr.f.
func Dormqr(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormqr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormrq.f.
func Sormrq(side blas.Side, trans blas.Transpose, m, n, k int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormrq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormrq.f.
func Dormrq(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormrq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormrz.f.
func Sormrz(side blas.Side, trans blas.Transpose, m, n, k, l int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormrz_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormrz.f.
func Dormrz(side blas.Side, trans blas.Transpose, m, n, k, l int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormrz_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sormtr.f.
func Sormtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, a []float32, lda int, tau, c []float32, ldc int, work []float32, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sormtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dormtr.f.
func Dormtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dormtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbcon.f.
func Spbcon(ul blas.Uplo, n, kd int, ab []float32, ldab int, anorm float32, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_spbcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbcon.f.
func Dpbcon(ul blas.Uplo, n, kd int, ab []float64, ldab int, anorm float64, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dpbcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbcon.f.
func Cpbcon(ul blas.Uplo, n, kd int, ab []complex64, ldab int, anorm float32, rcond []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cpbcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbcon.f.
func Zpbcon(ul blas.Uplo, n, kd int, ab []complex128, ldab int, anorm float64, rcond []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zpbcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbequ.f.
func Spbequ(ul blas.Uplo, n, kd int, ab []float32, ldab int, s, scond, amax []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_spbequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbequ.f.
func Dpbequ(ul blas.Uplo, n, kd int, ab []float64, ldab int, s, scond, amax []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dpbequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbequ.f.
func Cpbequ(ul blas.Uplo, n, kd int, ab []complex64, ldab int, s, scond, amax []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cpbequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbequ.f.
func Zpbequ(ul blas.Uplo, n, kd int, ab []complex128, ldab int, s, scond, amax []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zpbequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbrfs.f.
func Spbrfs(ul blas.Uplo, n, kd, nrhs int, ab []float32, ldab int, afb []float32, ldafb int, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float32
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_spbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_afb), (C.lapack_int)(ldafb), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbrfs.f.
func Dpbrfs(ul blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, afb []float64, ldafb int, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dpbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_afb), (C.lapack_int)(ldafb), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbrfs.f.
func Cpbrfs(ul blas.Uplo, n, kd, nrhs int, ab []complex64, ldab int, afb []complex64, ldafb int, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cpbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_afb), (C.lapack_int)(ldafb), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbrfs.f.
func Zpbrfs(ul blas.Uplo, n, kd, nrhs int, ab []complex128, ldab int, afb []complex128, ldafb int, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex128
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zpbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_afb), (C.lapack_int)(ldafb), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbstf.f.
func Spbstf(ul blas.Uplo, n, kb int, bb []float32, ldbb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _bb *float32
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	return isZero(C.LAPACKE_spbstf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kb), (*C.float)(_bb), (C.lapack_int)(ldbb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbstf.f.
func Dpbstf(ul blas.Uplo, n, kb int, bb []float64, ldbb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _bb *float64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	return isZero(C.LAPACKE_dpbstf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kb), (*C.double)(_bb), (C.lapack_int)(ldbb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbstf.f.
func Cpbstf(ul blas.Uplo, n, kb int, bb []complex64, ldbb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _bb *complex64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	return isZero(C.LAPACKE_cpbstf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kb), (*C.lapack_complex_float)(_bb), (C.lapack_int)(ldbb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbstf.f.
func Zpbstf(ul blas.Uplo, n, kb int, bb []complex128, ldbb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _bb *complex128
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	return isZero(C.LAPACKE_zpbstf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kb), (*C.lapack_complex_double)(_bb), (C.lapack_int)(ldbb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbsv.f.
func Spbsv(ul blas.Uplo, n, kd, nrhs int, ab []float32, ldab int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_spbsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbsv.f.
func Dpbsv(ul blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dpbsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbsv.f.
func Cpbsv(ul blas.Uplo, n, kd, nrhs int, ab []complex64, ldab int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cpbsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbsv.f.
func Zpbsv(ul blas.Uplo, n, kd, nrhs int, ab []complex128, ldab int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zpbsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbsvx.f.
func Spbsvx(fact byte, ul blas.Uplo, n, kd, nrhs int, ab []float32, ldab int, afb []float32, ldafb int, equed []byte, s, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float32
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_spbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_afb), (C.lapack_int)(ldafb), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_s), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbsvx.f.
func Dpbsvx(fact byte, ul blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, afb []float64, ldafb int, equed []byte, s, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *float64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dpbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_afb), (C.lapack_int)(ldafb), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_s), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbsvx.f.
func Cpbsvx(fact byte, ul blas.Uplo, n, kd, nrhs int, ab []complex64, ldab int, afb []complex64, ldafb int, equed []byte, s []float32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex64
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cpbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_afb), (C.lapack_int)(ldafb), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_s), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbsvx.f.
func Zpbsvx(fact byte, ul blas.Uplo, n, kd, nrhs int, ab []complex128, ldab int, afb []complex128, ldafb int, equed []byte, s []float64, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _afb *complex128
	if len(afb) > 0 {
		_afb = &afb[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zpbsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_afb), (C.lapack_int)(ldafb), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_s), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbtrf.f.
func Spbtrf(ul blas.Uplo, n, kd int, ab []float32, ldab int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	return isZero(C.LAPACKE_spbtrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbtrf.f.
func Dpbtrf(ul blas.Uplo, n, kd int, ab []float64, ldab int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	return isZero(C.LAPACKE_dpbtrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbtrf.f.
func Cpbtrf(ul blas.Uplo, n, kd int, ab []complex64, ldab int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	return isZero(C.LAPACKE_cpbtrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbtrf.f.
func Zpbtrf(ul blas.Uplo, n, kd int, ab []complex128, ldab int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	return isZero(C.LAPACKE_zpbtrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spbtrs.f.
func Spbtrs(ul blas.Uplo, n, kd, nrhs int, ab []float32, ldab int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_spbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpbtrs.f.
func Dpbtrs(ul blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dpbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpbtrs.f.
func Cpbtrs(ul blas.Uplo, n, kd, nrhs int, ab []complex64, ldab int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cpbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpbtrs.f.
func Zpbtrs(ul blas.Uplo, n, kd, nrhs int, ab []complex128, ldab int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zpbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spftrf.f.
func Spftrf(transr blas.Transpose, ul blas.Uplo, n int, a []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_spftrf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpftrf.f.
func Dpftrf(transr blas.Transpose, ul blas.Uplo, n int, a []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dpftrf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpftrf.f.
func Cpftrf(transr blas.Transpose, ul blas.Uplo, n int, a []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_cpftrf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpftrf.f.
func Zpftrf(transr blas.Transpose, ul blas.Uplo, n int, a []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zpftrf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spftri.f.
func Spftri(transr blas.Transpose, ul blas.Uplo, n int, a []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_spftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpftri.f.
func Dpftri(transr blas.Transpose, ul blas.Uplo, n int, a []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dpftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpftri.f.
func Cpftri(transr blas.Transpose, ul blas.Uplo, n int, a []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_cpftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpftri.f.
func Zpftri(transr blas.Transpose, ul blas.Uplo, n int, a []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zpftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spftrs.f.
func Spftrs(transr blas.Transpose, ul blas.Uplo, n, nrhs int, a, b []float32, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_spftrs_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpftrs.f.
func Dpftrs(transr blas.Transpose, ul blas.Uplo, n, nrhs int, a, b []float64, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dpftrs_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpftrs.f.
func Cpftrs(transr blas.Transpose, ul blas.Uplo, n, nrhs int, a, b []complex64, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cpftrs_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpftrs.f.
func Zpftrs(transr blas.Transpose, ul blas.Uplo, n, nrhs int, a, b []complex128, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zpftrs_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spocon.f.
func Spocon(ul blas.Uplo, n int, a []float32, lda int, anorm float32, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_spocon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpocon.f.
func Dpocon(ul blas.Uplo, n int, a []float64, lda int, anorm float64, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dpocon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpocon.f.
func Cpocon(ul blas.Uplo, n int, a []complex64, lda int, anorm float32, rcond []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cpocon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpocon.f.
func Zpocon(ul blas.Uplo, n int, a []complex128, lda int, anorm float64, rcond []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zpocon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spoequ.f.
func Spoequ(n int, a []float32, lda int, s, scond, amax []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_spoequ_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpoequ.f.
func Dpoequ(n int, a []float64, lda int, s, scond, amax []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dpoequ_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpoequ.f.
func Cpoequ(n int, a []complex64, lda int, s, scond, amax []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cpoequ_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpoequ.f.
func Zpoequ(n int, a []complex128, lda int, s, scond, amax []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zpoequ_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spoequb.f.
func Spoequb(n int, a []float32, lda int, s, scond, amax []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_spoequb_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpoequb.f.
func Dpoequb(n int, a []float64, lda int, s, scond, amax []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dpoequb_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpoequb.f.
func Cpoequb(n int, a []complex64, lda int, s, scond, amax []float32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cpoequb_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpoequb.f.
func Zpoequb(n int, a []complex128, lda int, s, scond, amax []float64) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zpoequb_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sporfs.f.
func Sporfs(ul blas.Uplo, n, nrhs int, a []float32, lda int, af []float32, ldaf int, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float32
	if len(af) > 0 {
		_af = &af[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sporfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_af), (C.lapack_int)(ldaf), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dporfs.f.
func Dporfs(ul blas.Uplo, n, nrhs int, a []float64, lda int, af []float64, ldaf int, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dporfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_af), (C.lapack_int)(ldaf), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cporfs.f.
func Cporfs(ul blas.Uplo, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cporfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zporfs.f.
func Zporfs(ul blas.Uplo, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zporfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sposv.f.
func Sposv(ul blas.Uplo, n, nrhs int, a []float32, lda int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sposv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dposv.f.
func Dposv(ul blas.Uplo, n, nrhs int, a []float64, lda int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dposv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cposv.f.
func Cposv(ul blas.Uplo, n, nrhs int, a []complex64, lda int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cposv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zposv.f.
func Zposv(ul blas.Uplo, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zposv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsposv.f.
func Dsposv(ul blas.Uplo, n, nrhs int, a []float64, lda int, b []float64, ldb int, x []float64, ldx int, work []float64, swork []float32, iter []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _swork *float32
	if len(swork) > 0 {
		_swork = &swork[0]
	}
	var _iter *int32
	if len(iter) > 0 {
		_iter = &iter[0]
	}
	return isZero(C.LAPACKE_dsposv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_work), (*C.float)(_swork), (*C.lapack_int)(_iter)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zcposv.f.
func Zcposv(ul blas.Uplo, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, x []complex128, ldx int, work []complex128, swork []complex64, rwork []float64, iter []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _swork *complex64
	if len(swork) > 0 {
		_swork = &swork[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iter *int32
	if len(iter) > 0 {
		_iter = &iter[0]
	}
	return isZero(C.LAPACKE_zcposv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.lapack_complex_double)(_work), (*C.lapack_complex_float)(_swork), (*C.double)(_rwork), (*C.lapack_int)(_iter)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sposvx.f.
func Sposvx(fact byte, ul blas.Uplo, n, nrhs int, a []float32, lda int, af []float32, ldaf int, equed []byte, s, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float32
	if len(af) > 0 {
		_af = &af[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sposvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_af), (C.lapack_int)(ldaf), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_s), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dposvx.f.
func Dposvx(fact byte, ul blas.Uplo, n, nrhs int, a []float64, lda int, af []float64, ldaf int, equed []byte, s, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dposvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_af), (C.lapack_int)(ldaf), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_s), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cposvx.f.
func Cposvx(fact byte, ul blas.Uplo, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, equed []byte, s []float32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cposvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_s), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zposvx.f.
func Zposvx(fact byte, ul blas.Uplo, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, equed []byte, s []float64, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zposvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_s), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spotrf2.f.
func Spotrf2(ul blas.Uplo, n int, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_spotrf2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpotrf2.f.
func Dpotrf2(ul blas.Uplo, n int, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dpotrf2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpotrf2.f.
func Cpotrf2(ul blas.Uplo, n int, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_cpotrf2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpotrf2.f.
func Zpotrf2(ul blas.Uplo, n int, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zpotrf2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spotrf.f.
func Spotrf(ul blas.Uplo, n int, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_spotrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpotrf.f.
func Dpotrf(ul blas.Uplo, n int, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dpotrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpotrf.f.
func Cpotrf(ul blas.Uplo, n int, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_cpotrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpotrf.f.
func Zpotrf(ul blas.Uplo, n int, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zpotrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spotri.f.
func Spotri(ul blas.Uplo, n int, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_spotri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpotri.f.
func Dpotri(ul blas.Uplo, n int, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dpotri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpotri.f.
func Cpotri(ul blas.Uplo, n int, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_cpotri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpotri.f.
func Zpotri(ul blas.Uplo, n int, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zpotri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spotrs.f.
func Spotrs(ul blas.Uplo, n, nrhs int, a []float32, lda int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_spotrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpotrs.f.
func Dpotrs(ul blas.Uplo, n, nrhs int, a []float64, lda int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dpotrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpotrs.f.
func Cpotrs(ul blas.Uplo, n, nrhs int, a []complex64, lda int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cpotrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpotrs.f.
func Zpotrs(ul blas.Uplo, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zpotrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sppcon.f.
func Sppcon(ul blas.Uplo, n int, ap []float32, anorm float32, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sppcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dppcon.f.
func Dppcon(ul blas.Uplo, n int, ap []float64, anorm float64, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dppcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cppcon.f.
func Cppcon(ul blas.Uplo, n int, ap []complex64, anorm float32, rcond []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cppcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zppcon.f.
func Zppcon(ul blas.Uplo, n int, ap []complex128, anorm float64, rcond []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zppcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sppequ.f.
func Sppequ(ul blas.Uplo, n int, ap, s, scond, amax []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_sppequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dppequ.f.
func Dppequ(ul blas.Uplo, n int, ap, s, scond, amax []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_dppequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cppequ.f.
func Cppequ(ul blas.Uplo, n int, ap []complex64, s, scond, amax []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_cppequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zppequ.f.
func Zppequ(ul blas.Uplo, n int, ap []complex128, s, scond, amax []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	return isZero(C.LAPACKE_zppequ_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spprfs.f.
func Spprfs(ul blas.Uplo, n, nrhs int, ap, afp, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float32
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_spprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_afp), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpprfs.f.
func Dpprfs(ul blas.Uplo, n, nrhs int, ap, afp, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dpprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_afp), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpprfs.f.
func Cpprfs(ul blas.Uplo, n, nrhs int, ap, afp, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cpprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_afp), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpprfs.f.
func Zpprfs(ul blas.Uplo, n, nrhs int, ap, afp, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex128
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zpprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_afp), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sppsv.f.
func Sppsv(ul blas.Uplo, n, nrhs int, ap, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sppsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dppsv.f.
func Dppsv(ul blas.Uplo, n, nrhs int, ap, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dppsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cppsv.f.
func Cppsv(ul blas.Uplo, n, nrhs int, ap, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cppsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zppsv.f.
func Zppsv(ul blas.Uplo, n, nrhs int, ap, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zppsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sppsvx.f.
func Sppsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []float32, equed []byte, s, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float32
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sppsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_afp), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_s), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dppsvx.f.
func Dppsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []float64, equed []byte, s, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dppsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_afp), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_s), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cppsvx.f.
func Cppsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []complex64, equed []byte, s []float32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cppsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_afp), (*C.char)(unsafe.Pointer(_equed)), (*C.float)(_s), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zppsvx.f.
func Zppsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []complex128, equed []byte, s []float64, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex128
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _equed *byte
	if len(equed) > 0 {
		_equed = &equed[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zppsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_afp), (*C.char)(unsafe.Pointer(_equed)), (*C.double)(_s), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spptrf.f.
func Spptrf(ul blas.Uplo, n int, ap []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_spptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpptrf.f.
func Dpptrf(ul blas.Uplo, n int, ap []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_dpptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpptrf.f.
func Cpptrf(ul blas.Uplo, n int, ap []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_cpptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpptrf.f.
func Zpptrf(ul blas.Uplo, n int, ap []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_zpptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spptri.f.
func Spptri(ul blas.Uplo, n int, ap []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_spptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpptri.f.
func Dpptri(ul blas.Uplo, n int, ap []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_dpptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpptri.f.
func Cpptri(ul blas.Uplo, n int, ap []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_cpptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpptri.f.
func Zpptri(ul blas.Uplo, n int, ap []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_zpptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spptrs.f.
func Spptrs(ul blas.Uplo, n, nrhs int, ap, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_spptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpptrs.f.
func Dpptrs(ul blas.Uplo, n, nrhs int, ap, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dpptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpptrs.f.
func Cpptrs(ul blas.Uplo, n, nrhs int, ap, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cpptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpptrs.f.
func Zpptrs(ul blas.Uplo, n, nrhs int, ap, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zpptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spstrf.f.
func Spstrf(ul blas.Uplo, n int, a []float32, lda int, piv, rank []int32, tol float32, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _piv *int32
	if len(piv) > 0 {
		_piv = &piv[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_spstrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_piv), (*C.lapack_int)(_rank), (C.float)(tol), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpstrf.f.
func Dpstrf(ul blas.Uplo, n int, a []float64, lda int, piv, rank []int32, tol float64, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _piv *int32
	if len(piv) > 0 {
		_piv = &piv[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dpstrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_piv), (*C.lapack_int)(_rank), (C.double)(tol), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpstrf.f.
func Cpstrf(ul blas.Uplo, n int, a []complex64, lda int, piv, rank []int32, tol float32, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _piv *int32
	if len(piv) > 0 {
		_piv = &piv[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cpstrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_piv), (*C.lapack_int)(_rank), (C.float)(tol), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpstrf.f.
func Zpstrf(ul blas.Uplo, n int, a []complex128, lda int, piv, rank []int32, tol float64, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _piv *int32
	if len(piv) > 0 {
		_piv = &piv[0]
	}
	var _rank *int32
	if len(rank) > 0 {
		_rank = &rank[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zpstrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_piv), (*C.lapack_int)(_rank), (C.double)(tol), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sptcon.f.
func Sptcon(n int, d, e []float32, anorm float32, rcond, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sptcon_work((C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dptcon.f.
func Dptcon(n int, d, e []float64, anorm float64, rcond, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dptcon_work((C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cptcon.f.
func Cptcon(n int, d []float32, e []complex64, anorm float32, rcond, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cptcon_work((C.lapack_int)(n), (*C.float)(_d), (*C.lapack_complex_float)(_e), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zptcon.f.
func Zptcon(n int, d []float64, e []complex128, anorm float64, rcond, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zptcon_work((C.lapack_int)(n), (*C.double)(_d), (*C.lapack_complex_double)(_e), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spteqr.f.
func Spteqr(compz lapack.Comp, n int, d, e, z []float32, ldz int, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_spteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpteqr.f.
func Dpteqr(compz lapack.Comp, n int, d, e, z []float64, ldz int, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dpteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpteqr.f.
func Cpteqr(compz lapack.Comp, n int, d, e []float32, z []complex64, ldz int, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cpteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpteqr.f.
func Zpteqr(compz lapack.Comp, n int, d, e []float64, z []complex128, ldz int, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zpteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sptrfs.f.
func Sptrfs(n, nrhs int, d, e, df, ef, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float32
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *float32
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sptrfs_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.float)(_e), (*C.float)(_df), (*C.float)(_ef), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dptrfs.f.
func Dptrfs(n, nrhs int, d, e, df, ef, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *float64
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dptrfs_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.double)(_e), (*C.double)(_df), (*C.double)(_ef), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cptrfs.f.
func Cptrfs(ul blas.Uplo, n, nrhs int, d []float32, e []complex64, df []float32, ef, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float32
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *complex64
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cptrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.lapack_complex_float)(_e), (*C.float)(_df), (*C.lapack_complex_float)(_ef), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zptrfs.f.
func Zptrfs(ul blas.Uplo, n, nrhs int, d []float64, e []complex128, df []float64, ef, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *complex128
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zptrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.lapack_complex_double)(_e), (*C.double)(_df), (*C.lapack_complex_double)(_ef), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sptsv.f.
func Sptsv(n, nrhs int, d, e, b []float32, ldb int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sptsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.float)(_e), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dptsv.f.
func Dptsv(n, nrhs int, d, e, b []float64, ldb int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dptsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.double)(_e), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cptsv.f.
func Cptsv(n, nrhs int, d []float32, e, b []complex64, ldb int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cptsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.lapack_complex_float)(_e), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zptsv.f.
func Zptsv(n, nrhs int, d []float64, e, b []complex128, ldb int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zptsv_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.lapack_complex_double)(_e), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sptsvx.f.
func Sptsvx(fact byte, n, nrhs int, d, e, df, ef, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float32
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *float32
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sptsvx_work((C.int)(rowMajor), (C.char)(fact), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.float)(_e), (*C.float)(_df), (*C.float)(_ef), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dptsvx.f.
func Dptsvx(fact byte, n, nrhs int, d, e, df, ef, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *float64
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dptsvx_work((C.int)(rowMajor), (C.char)(fact), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.double)(_e), (*C.double)(_df), (*C.double)(_ef), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cptsvx.f.
func Cptsvx(fact byte, n, nrhs int, d []float32, e []complex64, df []float32, ef, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float32
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *complex64
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cptsvx_work((C.int)(rowMajor), (C.char)(fact), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.lapack_complex_float)(_e), (*C.float)(_df), (*C.lapack_complex_float)(_ef), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zptsvx.f.
func Zptsvx(fact byte, n, nrhs int, d []float64, e []complex128, df []float64, ef, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	var _df *float64
	if len(df) > 0 {
		_df = &df[0]
	}
	var _ef *complex128
	if len(ef) > 0 {
		_ef = &ef[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zptsvx_work((C.int)(rowMajor), (C.char)(fact), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.lapack_complex_double)(_e), (*C.double)(_df), (*C.lapack_complex_double)(_ef), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spttrf.f.
func Spttrf(n int, d, e []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	return isZero(C.LAPACKE_spttrf_work((C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpttrf.f.
func Dpttrf(n int, d, e []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	return isZero(C.LAPACKE_dpttrf_work((C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpttrf.f.
func Cpttrf(n int, d []float32, e []complex64) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	return isZero(C.LAPACKE_cpttrf_work((C.lapack_int)(n), (*C.float)(_d), (*C.lapack_complex_float)(_e)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpttrf.f.
func Zpttrf(n int, d []float64, e []complex128) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	return isZero(C.LAPACKE_zpttrf_work((C.lapack_int)(n), (*C.double)(_d), (*C.lapack_complex_double)(_e)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/spttrs.f.
func Spttrs(n, nrhs int, d, e, b []float32, ldb int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_spttrs_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.float)(_e), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dpttrs.f.
func Dpttrs(n, nrhs int, d, e, b []float64, ldb int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dpttrs_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.double)(_e), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cpttrs.f.
func Cpttrs(ul blas.Uplo, n, nrhs int, d []float32, e, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cpttrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_d), (*C.lapack_complex_float)(_e), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zpttrs.f.
func Zpttrs(ul blas.Uplo, n, nrhs int, d []float64, e, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zpttrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_d), (*C.lapack_complex_double)(_e), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbev.f.
func Ssbev(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []float32, ldab int, w, z []float32, ldz int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssbev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbev.f.
func Dsbev(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []float64, ldab int, w, z []float64, ldz int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsbev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbevd.f.
func Ssbevd(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []float32, ldab int, w, z []float32, ldz int, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssbevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbevd.f.
func Dsbevd(jobz lapack.Job, ul blas.Uplo, n, kd int, ab []float64, ldab int, w, z []float64, ldz int, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsbevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbevx.f.
func Ssbevx(jobz lapack.Job, rng byte, ul blas.Uplo, n, kd int, ab []float32, ldab int, q []float32, ldq int, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_ssbevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_q), (C.lapack_int)(ldq), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbevx.f.
func Dsbevx(jobz lapack.Job, rng byte, ul blas.Uplo, n, kd int, ab []float64, ldab int, q []float64, ldq int, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dsbevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_q), (C.lapack_int)(ldq), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbgst.f.
func Ssbgst(vect byte, ul blas.Uplo, n, ka, kb int, ab []float32, ldab int, bb []float32, ldbb int, x []float32, ldx int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float32
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssbgst_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_bb), (C.lapack_int)(ldbb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbgst.f.
func Dsbgst(vect byte, ul blas.Uplo, n, ka, kb int, ab []float64, ldab int, bb []float64, ldbb int, x []float64, ldx int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsbgst_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_bb), (C.lapack_int)(ldbb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbgv.f.
func Ssbgv(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []float32, ldab int, bb []float32, ldbb int, w, z []float32, ldz int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float32
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssbgv_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_bb), (C.lapack_int)(ldbb), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbgv.f.
func Dsbgv(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []float64, ldab int, bb []float64, ldbb int, w, z []float64, ldz int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsbgv_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_bb), (C.lapack_int)(ldbb), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbgvd.f.
func Ssbgvd(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []float32, ldab int, bb []float32, ldbb int, w, z []float32, ldz int, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float32
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssbgvd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_bb), (C.lapack_int)(ldbb), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbgvd.f.
func Dsbgvd(jobz lapack.Job, ul blas.Uplo, n, ka, kb int, ab []float64, ldab int, bb []float64, ldbb int, w, z []float64, ldz int, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsbgvd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_bb), (C.lapack_int)(ldbb), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbgvx.f.
func Ssbgvx(jobz lapack.Job, rng byte, ul blas.Uplo, n, ka, kb int, ab []float32, ldab int, bb []float32, ldbb int, q []float32, ldq int, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float32
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_ssbgvx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_bb), (C.lapack_int)(ldbb), (*C.float)(_q), (C.lapack_int)(ldq), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbgvx.f.
func Dsbgvx(jobz lapack.Job, rng byte, ul blas.Uplo, n, ka, kb int, ab []float64, ldab int, bb []float64, ldbb int, q []float64, ldq int, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _bb *float64
	if len(bb) > 0 {
		_bb = &bb[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dsbgvx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(ka), (C.lapack_int)(kb), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_bb), (C.lapack_int)(ldbb), (*C.double)(_q), (C.lapack_int)(ldq), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssbtrd.f.
func Ssbtrd(vect byte, ul blas.Uplo, n, kd int, ab []float32, ldab int, d, e, q []float32, ldq int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssbtrd_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_d), (*C.float)(_e), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsbtrd.f.
func Dsbtrd(vect byte, ul blas.Uplo, n, kd int, ab []float64, ldab int, d, e, q []float64, ldq int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsbtrd_work((C.int)(rowMajor), (C.char)(vect), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_d), (*C.double)(_e), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssfrk.f.
func Ssfrk(transr blas.Transpose, ul blas.Uplo, trans blas.Transpose, n, k int, alpha float32, a []float32, lda int, beta float32, c []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	return isZero(C.LAPACKE_ssfrk_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(k), (C.float)(alpha), (*C.float)(_a), (C.lapack_int)(lda), (C.float)(beta), (*C.float)(_c)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsfrk.f.
func Dsfrk(transr blas.Transpose, ul blas.Uplo, trans blas.Transpose, n, k int, alpha float64, a []float64, lda int, beta float64, c []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	return isZero(C.LAPACKE_dsfrk_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(trans), (C.lapack_int)(n), (C.lapack_int)(k), (C.double)(alpha), (*C.double)(_a), (C.lapack_int)(lda), (C.double)(beta), (*C.double)(_c)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspcon.f.
func Sspcon(ul blas.Uplo, n int, ap []float32, ipiv []int32, anorm float32, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sspcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspcon.f.
func Dspcon(ul blas.Uplo, n int, ap []float64, ipiv []int32, anorm float64, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dspcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cspcon.f.
func Cspcon(ul blas.Uplo, n int, ap []complex64, ipiv []int32, anorm float32, rcond []float32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cspcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zspcon.f.
func Zspcon(ul blas.Uplo, n int, ap []complex128, ipiv []int32, anorm float64, rcond []float64, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zspcon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspev.f.
func Sspev(jobz lapack.Job, ul blas.Uplo, n int, ap, w, z []float32, ldz int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sspev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspev.f.
func Dspev(jobz lapack.Job, ul blas.Uplo, n int, ap, w, z []float64, ldz int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dspev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspevd.f.
func Sspevd(jobz lapack.Job, ul blas.Uplo, n int, ap, w, z []float32, ldz int, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sspevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspevd.f.
func Dspevd(jobz lapack.Job, ul blas.Uplo, n int, ap, w, z []float64, ldz int, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dspevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspevx.f.
func Sspevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap []float32, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_sspevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspevx.f.
func Dspevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap []float64, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dspevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspgst.f.
func Sspgst(itype int, ul blas.Uplo, n int, ap, bp []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float32
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	return isZero(C.LAPACKE_sspgst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_bp)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspgst.f.
func Dspgst(itype int, ul blas.Uplo, n int, ap, bp []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	return isZero(C.LAPACKE_dspgst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_bp)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspgv.f.
func Sspgv(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp, w, z []float32, ldz int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float32
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sspgv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_bp), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspgv.f.
func Dspgv(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp, w, z []float64, ldz int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dspgv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_bp), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspgvd.f.
func Sspgvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp, w, z []float32, ldz int, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float32
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sspgvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_bp), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspgvd.f.
func Dspgvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, ap, bp, w, z []float64, ldz int, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dspgvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_bp), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspgvx.f.
func Sspgvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap, bp []float32, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float32
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_sspgvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_bp), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspgvx.f.
func Dspgvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, ap, bp []float64, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _bp *float64
	if len(bp) > 0 {
		_bp = &bp[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dspgvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_bp), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssprfs.f.
func Ssprfs(ul blas.Uplo, n, nrhs int, ap, afp []float32, ipiv []int32, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float32
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_afp), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsprfs.f.
func Dsprfs(ul blas.Uplo, n, nrhs int, ap, afp []float64, ipiv []int32, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_afp), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csprfs.f.
func Csprfs(ul blas.Uplo, n, nrhs int, ap, afp []complex64, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_csprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsprfs.f.
func Zsprfs(ul blas.Uplo, n, nrhs int, ap, afp []complex128, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex128
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zsprfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspsv.f.
func Sspsv(ul blas.Uplo, n, nrhs int, ap []float32, ipiv []int32, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_sspsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspsv.f.
func Dspsv(ul blas.Uplo, n, nrhs int, ap []float64, ipiv []int32, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dspsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cspsv.f.
func Cspsv(ul blas.Uplo, n, nrhs int, ap []complex64, ipiv []int32, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_cspsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zspsv.f.
func Zspsv(ul blas.Uplo, n, nrhs int, ap []complex128, ipiv []int32, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zspsv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sspsvx.f.
func Sspsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []float32, ipiv []int32, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float32
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sspsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_afp), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dspsvx.f.
func Dspsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []float64, ipiv []int32, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *float64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dspsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_afp), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cspsvx.f.
func Cspsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []complex64, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex64
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cspsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zspsvx.f.
func Zspsvx(fact byte, ul blas.Uplo, n, nrhs int, ap, afp []complex128, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _afp *complex128
	if len(afp) > 0 {
		_afp = &afp[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zspsvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_afp), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssptrd.f.
func Ssptrd(ul blas.Uplo, n int, ap, d, e, tau []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_ssptrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_d), (*C.float)(_e), (*C.float)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsptrd.f.
func Dsptrd(ul blas.Uplo, n int, ap, d, e, tau []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	return isZero(C.LAPACKE_dsptrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_d), (*C.double)(_e), (*C.double)(_tau)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssptrf.f.
func Ssptrf(ul blas.Uplo, n int, ap []float32, ipiv []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_ssptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsptrf.f.
func Dsptrf(ul blas.Uplo, n int, ap []float64, ipiv []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_dsptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csptrf.f.
func Csptrf(ul blas.Uplo, n int, ap []complex64, ipiv []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_csptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsptrf.f.
func Zsptrf(ul blas.Uplo, n int, ap []complex128, ipiv []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	return isZero(C.LAPACKE_zsptrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssptri.f.
func Ssptri(ul blas.Uplo, n int, ap []float32, ipiv []int32, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.lapack_int)(_ipiv), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsptri.f.
func Dsptri(ul blas.Uplo, n int, ap []float64, ipiv []int32, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.lapack_int)(_ipiv), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csptri.f.
func Csptri(ul blas.Uplo, n int, ap []complex64, ipiv []int32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsptri.f.
func Zsptri(ul blas.Uplo, n int, ap []complex128, ipiv []int32, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsptri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssptrs.f.
func Ssptrs(ul blas.Uplo, n, nrhs int, ap []float32, ipiv []int32, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ssptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsptrs.f.
func Dsptrs(ul blas.Uplo, n, nrhs int, ap []float64, ipiv []int32, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dsptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csptrs.f.
func Csptrs(ul blas.Uplo, n, nrhs int, ap []complex64, ipiv []int32, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_csptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsptrs.f.
func Zsptrs(ul blas.Uplo, n, nrhs int, ap []complex128, ipiv []int32, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zsptrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstebz.f.
func Sstebz(rng, order byte, n int, vl, vu float32, il, iu int, abstol float32, d, e []float32, m, nsplit []int32, w []float32, iblock, isplit []int32, work []float32, iwork []int32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _nsplit *int32
	if len(nsplit) > 0 {
		_nsplit = &nsplit[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _iblock *int32
	if len(iblock) > 0 {
		_iblock = &iblock[0]
	}
	var _isplit *int32
	if len(isplit) > 0 {
		_isplit = &isplit[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sstebz_work((C.char)(rng), (C.char)(order), (C.lapack_int)(n), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.float)(_d), (*C.float)(_e), (*C.lapack_int)(_m), (*C.lapack_int)(_nsplit), (*C.float)(_w), (*C.lapack_int)(_iblock), (*C.lapack_int)(_isplit), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstebz.f.
func Dstebz(rng, order byte, n int, vl, vu float64, il, iu int, abstol float64, d, e []float64, m, nsplit []int32, w []float64, iblock, isplit []int32, work []float64, iwork []int32) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _nsplit *int32
	if len(nsplit) > 0 {
		_nsplit = &nsplit[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _iblock *int32
	if len(iblock) > 0 {
		_iblock = &iblock[0]
	}
	var _isplit *int32
	if len(isplit) > 0 {
		_isplit = &isplit[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dstebz_work((C.char)(rng), (C.char)(order), (C.lapack_int)(n), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.double)(_d), (*C.double)(_e), (*C.lapack_int)(_m), (*C.lapack_int)(_nsplit), (*C.double)(_w), (*C.lapack_int)(_iblock), (*C.lapack_int)(_isplit), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstedc.f.
func Sstedc(compz lapack.Comp, n int, d, e, z []float32, ldz int, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sstedc_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstedc.f.
func Dstedc(compz lapack.Comp, n int, d, e, z []float64, ldz int, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dstedc_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cstedc.f.
func Cstedc(compz lapack.Comp, n int, d, e []float32, z []complex64, ldz int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cstedc_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zstedc.f.
func Zstedc(compz lapack.Comp, n int, d, e []float64, z []complex128, ldz int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zstedc_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstegr.f.
func Sstegr(jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, isuppz []int32, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sstegr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstegr.f.
func Dstegr(jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, isuppz []int32, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dstegr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cstegr.f.
func Cstegr(jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu float32, il, iu int, abstol float32, m []int32, w []float32, z []complex64, ldz int, isuppz []int32, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cstegr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zstegr.f.
func Zstegr(jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu float64, il, iu int, abstol float64, m []int32, w []float64, z []complex128, ldz int, isuppz []int32, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zstegr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstein.f.
func Sstein(n int, d, e []float32, m int, w []float32, iblock, isplit []int32, z []float32, ldz int, work []float32, iwork, ifailv []int32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _iblock *int32
	if len(iblock) > 0 {
		_iblock = &iblock[0]
	}
	var _isplit *int32
	if len(isplit) > 0 {
		_isplit = &isplit[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifailv *int32
	if len(ifailv) > 0 {
		_ifailv = &ifailv[0]
	}
	return isZero(C.LAPACKE_sstein_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.lapack_int)(m), (*C.float)(_w), (*C.lapack_int)(_iblock), (*C.lapack_int)(_isplit), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifailv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstein.f.
func Dstein(n int, d, e []float64, m int, w []float64, iblock, isplit []int32, z []float64, ldz int, work []float64, iwork, ifailv []int32) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _iblock *int32
	if len(iblock) > 0 {
		_iblock = &iblock[0]
	}
	var _isplit *int32
	if len(isplit) > 0 {
		_isplit = &isplit[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifailv *int32
	if len(ifailv) > 0 {
		_ifailv = &ifailv[0]
	}
	return isZero(C.LAPACKE_dstein_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.lapack_int)(m), (*C.double)(_w), (*C.lapack_int)(_iblock), (*C.lapack_int)(_isplit), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifailv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cstein.f.
func Cstein(n int, d, e []float32, m int, w []float32, iblock, isplit []int32, z []complex64, ldz int, work []float32, iwork, ifailv []int32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _iblock *int32
	if len(iblock) > 0 {
		_iblock = &iblock[0]
	}
	var _isplit *int32
	if len(isplit) > 0 {
		_isplit = &isplit[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifailv *int32
	if len(ifailv) > 0 {
		_ifailv = &ifailv[0]
	}
	return isZero(C.LAPACKE_cstein_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.lapack_int)(m), (*C.float)(_w), (*C.lapack_int)(_iblock), (*C.lapack_int)(_isplit), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifailv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zstein.f.
func Zstein(n int, d, e []float64, m int, w []float64, iblock, isplit []int32, z []complex128, ldz int, work []float64, iwork, ifailv []int32) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _iblock *int32
	if len(iblock) > 0 {
		_iblock = &iblock[0]
	}
	var _isplit *int32
	if len(isplit) > 0 {
		_isplit = &isplit[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifailv *int32
	if len(ifailv) > 0 {
		_ifailv = &ifailv[0]
	}
	return isZero(C.LAPACKE_zstein_work((C.int)(rowMajor), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.lapack_int)(m), (*C.double)(_w), (*C.lapack_int)(_iblock), (*C.lapack_int)(_isplit), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifailv)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstemr.f.
func Sstemr(jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu float32, il, iu int, m []int32, w, z []float32, ldz, nzc int, isuppz, tryrac []int32, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _tryrac *int32
	if len(tryrac) > 0 {
		_tryrac = &tryrac[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sstemr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (C.lapack_int)(nzc), (*C.lapack_int)(_isuppz), (*C.lapack_logical)(_tryrac), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstemr.f.
func Dstemr(jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu float64, il, iu int, m []int32, w, z []float64, ldz, nzc int, isuppz, tryrac []int32, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _tryrac *int32
	if len(tryrac) > 0 {
		_tryrac = &tryrac[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dstemr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (C.lapack_int)(nzc), (*C.lapack_int)(_isuppz), (*C.lapack_logical)(_tryrac), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cstemr.f.
func Cstemr(jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu float32, il, iu int, m []int32, w []float32, z []complex64, ldz, nzc int, isuppz, tryrac []int32, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _tryrac *int32
	if len(tryrac) > 0 {
		_tryrac = &tryrac[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cstemr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (*C.lapack_int)(_m), (*C.float)(_w), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (C.lapack_int)(nzc), (*C.lapack_int)(_isuppz), (*C.lapack_logical)(_tryrac), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zstemr.f.
func Zstemr(jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu float64, il, iu int, m []int32, w []float64, z []complex128, ldz, nzc int, isuppz, tryrac []int32, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _tryrac *int32
	if len(tryrac) > 0 {
		_tryrac = &tryrac[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zstemr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (*C.lapack_int)(_m), (*C.double)(_w), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (C.lapack_int)(nzc), (*C.lapack_int)(_isuppz), (*C.lapack_logical)(_tryrac), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssteqr.f.
func Ssteqr(compz lapack.Comp, n int, d, e, z []float32, ldz int, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsteqr.f.
func Dsteqr(compz lapack.Comp, n int, d, e, z []float64, ldz int, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csteqr.f.
func Csteqr(compz lapack.Comp, n int, d, e []float32, z []complex64, ldz int, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsteqr.f.
func Zsteqr(compz lapack.Comp, n int, d, e []float64, z []complex128, ldz int, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsteqr_work((C.int)(rowMajor), (C.char)(compz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssterf.f.
func Ssterf(n int, d, e []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	return isZero(C.LAPACKE_ssterf_work((C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsterf.f.
func Dsterf(n int, d, e []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	return isZero(C.LAPACKE_dsterf_work((C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstev.f.
func Sstev(jobz lapack.Job, n int, d, e, z []float32, ldz int, work []float32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sstev_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstev.f.
func Dstev(jobz lapack.Job, n int, d, e, z []float64, ldz int, work []float64) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dstev_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstevd.f.
func Sstevd(jobz lapack.Job, n int, d, e, z []float32, ldz int, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sstevd_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstevd.f.
func Dstevd(jobz lapack.Job, n int, d, e, z []float64, ldz int, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dstevd_work((C.int)(rowMajor), (C.char)(jobz), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstevr.f.
func Sstevr(jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, isuppz []int32, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sstevr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstevr.f.
func Dstevr(jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, isuppz []int32, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dstevr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sstevx.f.
func Sstevx(jobz lapack.Job, rng byte, n int, d, e []float32, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, iwork, ifail []int32) bool {
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_sstevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.float)(_d), (*C.float)(_e), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dstevx.f.
func Dstevx(jobz lapack.Job, rng byte, n int, d, e []float64, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, iwork, ifail []int32) bool {
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dstevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.lapack_int)(n), (*C.double)(_d), (*C.double)(_e), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssycon.f.
func Ssycon(ul blas.Uplo, n int, a []float32, lda int, ipiv []int32, anorm float32, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssycon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsycon.f.
func Dsycon(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32, anorm float64, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsycon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csycon.f.
func Csycon(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, anorm float32, rcond []float32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csycon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (C.float)(anorm), (*C.float)(_rcond), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsycon.f.
func Zsycon(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, anorm float64, rcond []float64, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsycon_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (C.double)(anorm), (*C.double)(_rcond), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyequb.f.
func Ssyequb(ul blas.Uplo, n int, a []float32, lda int, s, scond, amax, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssyequb_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyequb.f.
func Dsyequb(ul blas.Uplo, n int, a []float64, lda int, s, scond, amax, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsyequb_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csyequb.f.
func Csyequb(ul blas.Uplo, n int, a []complex64, lda int, s, scond, amax []float32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float32
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float32
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csyequb_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_s), (*C.float)(_scond), (*C.float)(_amax), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsyequb.f.
func Zsyequb(ul blas.Uplo, n int, a []complex128, lda int, s, scond, amax []float64, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _scond *float64
	if len(scond) > 0 {
		_scond = &scond[0]
	}
	var _amax *float64
	if len(amax) > 0 {
		_amax = &amax[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsyequb_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_s), (*C.double)(_scond), (*C.double)(_amax), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyev.f.
func Ssyev(jobz lapack.Job, ul blas.Uplo, n int, a []float32, lda int, w, work []float32, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssyev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_w), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyev.f.
func Dsyev(jobz lapack.Job, ul blas.Uplo, n int, a []float64, lda int, w, work []float64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsyev_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_w), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyevd.f.
func Ssyevd(jobz lapack.Job, ul blas.Uplo, n int, a []float32, lda int, w, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssyevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_w), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyevd.f.
func Dsyevd(jobz lapack.Job, ul blas.Uplo, n int, a []float64, lda int, w, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsyevd_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_w), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyevr.f.
func Ssyevr(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []float32, lda int, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, isuppz []int32, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssyevr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyevr.f.
func Dsyevr(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []float64, lda int, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, isuppz []int32, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _isuppz *int32
	if len(isuppz) > 0 {
		_isuppz = &isuppz[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsyevr_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_isuppz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyevx.f.
func Ssyevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []float32, lda int, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, lwork int, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_ssyevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyevx.f.
func Dsyevx(jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []float64, lda int, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, lwork int, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dsyevx_work((C.int)(rowMajor), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssygst.f.
func Ssygst(itype int, ul blas.Uplo, n int, a []float32, lda int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ssygst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsygst.f.
func Dsygst(itype int, ul blas.Uplo, n int, a []float64, lda int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dsygst_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssygv.f.
func Ssygv(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []float32, lda int, b []float32, ldb int, w, work []float32, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssygv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_w), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsygv.f.
func Dsygv(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []float64, lda int, b []float64, ldb int, w, work []float64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsygv_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_w), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssygvd.f.
func Ssygvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []float32, lda int, b []float32, ldb int, w, work []float32, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssygvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_w), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsygvd.f.
func Dsygvd(itype int, jobz lapack.Job, ul blas.Uplo, n int, a []float64, lda int, b []float64, ldb int, w, work []float64, lwork int, iwork []int32, liwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsygvd_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_w), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssygvx.f.
func Ssygvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []float32, lda int, b []float32, ldb int, vl, vu float32, il, iu int, abstol float32, m []int32, w, z []float32, ldz int, work []float32, lwork int, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float32
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_ssygvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (C.float)(vl), (C.float)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.float)(abstol), (*C.lapack_int)(_m), (*C.float)(_w), (*C.float)(_z), (C.lapack_int)(ldz), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsygvx.f.
func Dsygvx(itype int, jobz lapack.Job, rng byte, ul blas.Uplo, n int, a []float64, lda int, b []float64, ldb int, vl, vu float64, il, iu int, abstol float64, m []int32, w, z []float64, ldz int, work []float64, lwork int, iwork, ifail []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _w *float64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	var _ifail *int32
	if len(ifail) > 0 {
		_ifail = &ifail[0]
	}
	return isZero(C.LAPACKE_dsygvx_work((C.int)(rowMajor), (C.lapack_int)(itype), (C.char)(jobz), (C.char)(rng), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (C.double)(vl), (C.double)(vu), (C.lapack_int)(il), (C.lapack_int)(iu), (C.double)(abstol), (*C.lapack_int)(_m), (*C.double)(_w), (*C.double)(_z), (C.lapack_int)(ldz), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (*C.lapack_int)(_ifail)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyrfs.f.
func Ssyrfs(ul blas.Uplo, n, nrhs int, a []float32, lda int, af []float32, ldaf int, ipiv []int32, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float32
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssyrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyrfs.f.
func Dsyrfs(ul blas.Uplo, n, nrhs int, a []float64, lda int, af []float64, ldaf int, ipiv []int32, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsyrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csyrfs.f.
func Csyrfs(ul blas.Uplo, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_csyrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsyrfs.f.
func Zsyrfs(ul blas.Uplo, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zsyrfs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssysv.f.
func Ssysv(ul blas.Uplo, n, nrhs int, a []float32, lda int, ipiv []int32, b []float32, ldb int, work []float32, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssysv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsysv.f.
func Dsysv(ul blas.Uplo, n, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int, work []float64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsysv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csysv.f.
func Csysv(ul blas.Uplo, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csysv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsysv.f.
func Zsysv(ul blas.Uplo, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsysv_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssysvx.f.
func Ssysvx(fact byte, ul blas.Uplo, n, nrhs int, a []float32, lda int, af []float32, ldaf int, ipiv []int32, b []float32, ldb int, x []float32, ldx int, rcond, ferr, berr, work []float32, lwork int, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float32
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ssysvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsysvx.f.
func Dsysvx(fact byte, ul blas.Uplo, n, nrhs int, a []float64, lda int, af []float64, ldaf int, ipiv []int32, b []float64, ldb int, x []float64, ldx int, rcond, ferr, berr, work []float64, lwork int, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *float64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dsysvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csysvx.f.
func Csysvx(fact byte, ul blas.Uplo, n, nrhs int, a []complex64, lda int, af []complex64, ldaf int, ipiv []int32, b []complex64, ldb int, x []complex64, ldx int, rcond, ferr, berr []float32, work []complex64, lwork int, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex64
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_csysvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_rcond), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsysvx.f.
func Zsysvx(fact byte, ul blas.Uplo, n, nrhs int, a []complex128, lda int, af []complex128, ldaf int, ipiv []int32, b []complex128, ldb int, x []complex128, ldx int, rcond, ferr, berr []float64, work []complex128, lwork int, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _af *complex128
	if len(af) > 0 {
		_af = &af[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zsysvx_work((C.int)(rowMajor), (C.char)(fact), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_af), (C.lapack_int)(ldaf), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_rcond), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytrd.f.
func Ssytrd(ul blas.Uplo, n int, a []float32, lda int, d, e, tau, work []float32, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssytrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_d), (*C.float)(_e), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytrd.f.
func Dsytrd(ul blas.Uplo, n int, a []float64, lda int, d, e, tau, work []float64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsytrd_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_d), (*C.double)(_e), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytrf.f.
func Ssytrf(ul blas.Uplo, n int, a []float32, lda int, ipiv []int32, work []float32, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssytrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytrf.f.
func Dsytrf(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32, work []float64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsytrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csytrf.f.
func Csytrf(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csytrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsytrf.f.
func Zsytrf(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsytrf_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytri.f.
func Ssytri(ul blas.Uplo, n int, a []float32, lda int, ipiv []int32, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssytri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytri.f.
func Dsytri(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsytri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csytri.f.
func Csytri(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csytri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsytri.f.
func Zsytri(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsytri_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytrs.f.
func Ssytrs(ul blas.Uplo, n, nrhs int, a []float32, lda int, ipiv []int32, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ssytrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytrs.f.
func Dsytrs(ul blas.Uplo, n, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dsytrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csytrs.f.
func Csytrs(ul blas.Uplo, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_csytrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsytrs.f.
func Zsytrs(ul blas.Uplo, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_zsytrs_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stbcon.f.
func Stbcon(norm byte, ul blas.Uplo, d blas.Diag, n, kd int, ab []float32, ldab int, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stbcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtbcon.f.
func Dtbcon(norm byte, ul blas.Uplo, d blas.Diag, n, kd int, ab []float64, ldab int, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtbcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctbcon.f.
func Ctbcon(norm byte, ul blas.Uplo, d blas.Diag, n, kd int, ab []complex64, ldab int, rcond []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctbcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztbcon.f.
func Ztbcon(norm byte, ul blas.Uplo, d blas.Diag, n, kd int, ab []complex128, ldab int, rcond []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztbcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stbrfs.f.
func Stbrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []float32, ldab int, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtbrfs.f.
func Dtbrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctbrfs.f.
func Ctbrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []complex64, ldab int, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztbrfs.f.
func Ztbrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []complex128, ldab int, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztbrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stbtrs.f.
func Stbtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []float32, ldab int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *float32
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_stbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.float)(_ab), (C.lapack_int)(ldab), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtbtrs.f.
func Dtbtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *float64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dtbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.double)(_ab), (C.lapack_int)(ldab), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctbtrs.f.
func Ctbtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []complex64, ldab int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *complex64
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ctbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztbtrs.f.
func Ztbtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, kd, nrhs int, ab []complex128, ldab int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ab *complex128
	if len(ab) > 0 {
		_ab = &ab[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ztbtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(kd), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ab), (C.lapack_int)(ldab), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stfsm.f.
func Stfsm(transr blas.Transpose, side blas.Side, ul blas.Uplo, trans blas.Transpose, d blas.Diag, m, n int, alpha float32, a, b []float32, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_stfsm_work((C.int)(rowMajor), (C.char)(transr), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (C.float)(alpha), (*C.float)(_a), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtfsm.f.
func Dtfsm(transr blas.Transpose, side blas.Side, ul blas.Uplo, trans blas.Transpose, d blas.Diag, m, n int, alpha float64, a, b []float64, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dtfsm_work((C.int)(rowMajor), (C.char)(transr), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (C.double)(alpha), (*C.double)(_a), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctfsm.f.
func Ctfsm(transr blas.Transpose, side blas.Side, ul blas.Uplo, trans blas.Transpose, d blas.Diag, m, n int, alpha complex64, a, b []complex64, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ctfsm_work((C.int)(rowMajor), (C.char)(transr), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_complex_float)(alpha), (*C.lapack_complex_float)(_a), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztfsm.f.
func Ztfsm(transr blas.Transpose, side blas.Side, ul blas.Uplo, trans blas.Transpose, d blas.Diag, m, n int, alpha complex128, a, b []complex128, ldb int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ztfsm_work((C.int)(rowMajor), (C.char)(transr), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_complex_double)(alpha), (*C.lapack_complex_double)(_a), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stftri.f.
func Stftri(transr blas.Transpose, ul blas.Uplo, d blas.Diag, n int, a []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_stftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.float)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtftri.f.
func Dtftri(transr blas.Transpose, ul blas.Uplo, d blas.Diag, n int, a []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dtftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.double)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctftri.f.
func Ctftri(transr blas.Transpose, ul blas.Uplo, d blas.Diag, n int, a []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ctftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_float)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztftri.f.
func Ztftri(transr blas.Transpose, ul blas.Uplo, d blas.Diag, n int, a []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ztftri_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_double)(_a)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stfttp.f.
func Stfttp(transr blas.Transpose, ul blas.Uplo, n int, arf, ap []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *float32
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_stfttp_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_arf), (*C.float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtfttp.f.
func Dtfttp(transr blas.Transpose, ul blas.Uplo, n int, arf, ap []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *float64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_dtfttp_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_arf), (*C.double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctfttp.f.
func Ctfttp(transr blas.Transpose, ul blas.Uplo, n int, arf, ap []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *complex64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_ctfttp_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_arf), (*C.lapack_complex_float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztfttp.f.
func Ztfttp(transr blas.Transpose, ul blas.Uplo, n int, arf, ap []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *complex128
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_ztfttp_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_arf), (*C.lapack_complex_double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stfttr.f.
func Stfttr(transr blas.Transpose, ul blas.Uplo, n int, arf, a []float32, lda int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *float32
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_stfttr_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_arf), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtfttr.f.
func Dtfttr(transr blas.Transpose, ul blas.Uplo, n int, arf, a []float64, lda int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *float64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dtfttr_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_arf), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctfttr.f.
func Ctfttr(transr blas.Transpose, ul blas.Uplo, n int, arf, a []complex64, lda int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *complex64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ctfttr_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_arf), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztfttr.f.
func Ztfttr(transr blas.Transpose, ul blas.Uplo, n int, arf, a []complex128, lda int) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _arf *complex128
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ztfttr_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_arf), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stgevc.f.
func Stgevc(side blas.Side, howmny byte, sel []int32, n int, s []float32, lds int, p []float32, ldp int, vl []float32, ldvl int, vr []float32, ldvr, mm int, m []int32, work []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _p *float32
	if len(p) > 0 {
		_p = &p[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_stgevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_s), (C.lapack_int)(lds), (*C.float)(_p), (C.lapack_int)(ldp), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtgevc.f.
func Dtgevc(side blas.Side, howmny byte, sel []int32, n int, s []float64, lds int, p []float64, ldp int, vl []float64, ldvl int, vr []float64, ldvr, mm int, m []int32, work []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _p *float64
	if len(p) > 0 {
		_p = &p[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtgevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_s), (C.lapack_int)(lds), (*C.double)(_p), (C.lapack_int)(ldp), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctgevc.f.
func Ctgevc(side blas.Side, howmny byte, sel []int32, n int, s []complex64, lds int, p []complex64, ldp int, vl []complex64, ldvl int, vr []complex64, ldvr, mm int, m []int32, work []complex64, rwork []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _s *complex64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _p *complex64
	if len(p) > 0 {
		_p = &p[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctgevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_s), (C.lapack_int)(lds), (*C.lapack_complex_float)(_p), (C.lapack_int)(ldp), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztgevc.f.
func Ztgevc(side blas.Side, howmny byte, sel []int32, n int, s []complex128, lds int, p []complex128, ldp int, vl []complex128, ldvl int, vr []complex128, ldvr, mm int, m []int32, work []complex128, rwork []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _s *complex128
	if len(s) > 0 {
		_s = &s[0]
	}
	var _p *complex128
	if len(p) > 0 {
		_p = &p[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztgevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_s), (C.lapack_int)(lds), (*C.lapack_complex_double)(_p), (C.lapack_int)(ldp), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stgexc.f.
func Stgexc(wantq, wantz int32, n int, a []float32, lda int, b []float32, ldb int, q []float32, ldq int, z []float32, ldz int, ifst, ilst []int32, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _ifst *int32
	if len(ifst) > 0 {
		_ifst = &ifst[0]
	}
	var _ilst *int32
	if len(ilst) > 0 {
		_ilst = &ilst[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_stgexc_work((C.int)(rowMajor), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_ifst), (*C.lapack_int)(_ilst), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtgexc.f.
func Dtgexc(wantq, wantz int32, n int, a []float64, lda int, b []float64, ldb int, q []float64, ldq int, z []float64, ldz int, ifst, ilst []int32, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _ifst *int32
	if len(ifst) > 0 {
		_ifst = &ifst[0]
	}
	var _ilst *int32
	if len(ilst) > 0 {
		_ilst = &ilst[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtgexc_work((C.int)(rowMajor), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_ifst), (*C.lapack_int)(_ilst), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctgexc.f.
func Ctgexc(wantq, wantz int32, n int, a []complex64, lda int, b []complex64, ldb int, q []complex64, ldq int, z []complex64, ldz, ifst, ilst int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	return isZero(C.LAPACKE_ctgexc_work((C.int)(rowMajor), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (C.lapack_int)(ifst), (C.lapack_int)(ilst)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztgexc.f.
func Ztgexc(wantq, wantz int32, n int, a []complex128, lda int, b []complex128, ldb int, q []complex128, ldq int, z []complex128, ldz, ifst, ilst int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	return isZero(C.LAPACKE_ztgexc_work((C.int)(rowMajor), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (C.lapack_int)(ifst), (C.lapack_int)(ilst)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stgsen.f.
func Stgsen(ijob lapack.Job, wantq, wantz int32, sel []int32, n int, a []float32, lda int, b []float32, ldb int, alphar, alphai, beta, q []float32, ldq int, z []float32, ldz int, m []int32, pl, pr, dif, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float32
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float32
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float32
	if len(z) > 0 {
		_z = &z[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _pl *float32
	if len(pl) > 0 {
		_pl = &pl[0]
	}
	var _pr *float32
	if len(pr) > 0 {
		_pr = &pr[0]
	}
	var _dif *float32
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stgsen_work((C.int)(rowMajor), (C.lapack_int)(ijob), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_alphar), (*C.float)(_alphai), (*C.float)(_beta), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_m), (*C.float)(_pl), (*C.float)(_pr), (*C.float)(_dif), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtgsen.f.
func Dtgsen(ijob lapack.Job, wantq, wantz int32, sel []int32, n int, a []float64, lda int, b []float64, ldb int, alphar, alphai, beta, q []float64, ldq int, z []float64, ldz int, m []int32, pl, pr, dif, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alphar *float64
	if len(alphar) > 0 {
		_alphar = &alphar[0]
	}
	var _alphai *float64
	if len(alphai) > 0 {
		_alphai = &alphai[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *float64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _pl *float64
	if len(pl) > 0 {
		_pl = &pl[0]
	}
	var _pr *float64
	if len(pr) > 0 {
		_pr = &pr[0]
	}
	var _dif *float64
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtgsen_work((C.int)(rowMajor), (C.lapack_int)(ijob), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_alphar), (*C.double)(_alphai), (*C.double)(_beta), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_m), (*C.double)(_pl), (*C.double)(_pr), (*C.double)(_dif), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctgsen.f.
func Ctgsen(ijob lapack.Job, wantq, wantz int32, sel []int32, n int, a []complex64, lda int, b []complex64, ldb int, alpha, beta, q []complex64, ldq int, z []complex64, ldz int, m []int32, pl, pr, dif []float32, work []complex64, lwork int, iwork []int32, liwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex64
	if len(z) > 0 {
		_z = &z[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _pl *float32
	if len(pl) > 0 {
		_pl = &pl[0]
	}
	var _pr *float32
	if len(pr) > 0 {
		_pr = &pr[0]
	}
	var _dif *float32
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ctgsen_work((C.int)(rowMajor), (C.lapack_int)(ijob), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_alpha), (*C.lapack_complex_float)(_beta), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_m), (*C.float)(_pl), (*C.float)(_pr), (*C.float)(_dif), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztgsen.f.
func Ztgsen(ijob lapack.Job, wantq, wantz int32, sel []int32, n int, a []complex128, lda int, b []complex128, ldb int, alpha, beta, q []complex128, ldq int, z []complex128, ldz int, m []int32, pl, pr, dif []float64, work []complex128, lwork int, iwork []int32, liwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *complex128
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *complex128
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _z *complex128
	if len(z) > 0 {
		_z = &z[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _pl *float64
	if len(pl) > 0 {
		_pl = &pl[0]
	}
	var _pr *float64
	if len(pr) > 0 {
		_pr = &pr[0]
	}
	var _dif *float64
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ztgsen_work((C.int)(rowMajor), (C.lapack_int)(ijob), (C.lapack_logical)(wantq), (C.lapack_logical)(wantz), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_alpha), (*C.lapack_complex_double)(_beta), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_z), (C.lapack_int)(ldz), (*C.lapack_int)(_m), (*C.double)(_pl), (*C.double)(_pr), (*C.double)(_dif), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stgsja.f.
func Stgsja(jobu, jobv, jobq lapack.Job, m, p, n, k, l int, a []float32, lda int, b []float32, ldb int, tola, tolb float32, alpha, beta, u []float32, ldu int, v []float32, ldv int, q []float32, ldq int, work []float32, ncycle []int32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float32
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *float32
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _ncycle *int32
	if len(ncycle) > 0 {
		_ncycle = &ncycle[0]
	}
	return isZero(C.LAPACKE_stgsja_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (C.float)(tola), (C.float)(tolb), (*C.float)(_alpha), (*C.float)(_beta), (*C.float)(_u), (C.lapack_int)(ldu), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_work), (*C.lapack_int)(_ncycle)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtgsja.f.
func Dtgsja(jobu, jobv, jobq lapack.Job, m, p, n, k, l int, a []float64, lda int, b []float64, ldb int, tola, tolb float64, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64, ncycle []int32) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *float64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _ncycle *int32
	if len(ncycle) > 0 {
		_ncycle = &ncycle[0]
	}
	return isZero(C.LAPACKE_dtgsja_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (C.double)(tola), (C.double)(tolb), (*C.double)(_alpha), (*C.double)(_beta), (*C.double)(_u), (C.lapack_int)(ldu), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_work), (*C.lapack_int)(_ncycle)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctgsja.f.
func Ctgsja(jobu, jobv, jobq lapack.Job, m, p, n, k, l int, a []complex64, lda int, b []complex64, ldb int, tola, tolb float32, alpha, beta []float32, u []complex64, ldu int, v []complex64, ldv int, q []complex64, ldq int, work []complex64, ncycle []int32) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float32
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float32
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *complex64
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _ncycle *int32
	if len(ncycle) > 0 {
		_ncycle = &ncycle[0]
	}
	return isZero(C.LAPACKE_ctgsja_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (C.float)(tola), (C.float)(tolb), (*C.float)(_alpha), (*C.float)(_beta), (*C.lapack_complex_float)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_work), (*C.lapack_int)(_ncycle)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztgsja.f.
func Ztgsja(jobu, jobv, jobq lapack.Job, m, p, n, k, l int, a []complex128, lda int, b []complex128, ldb int, tola, tolb float64, alpha, beta []float64, u []complex128, ldu int, v []complex128, ldv int, q []complex128, ldq int, work []complex128, ncycle []int32) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _alpha *float64
	if len(alpha) > 0 {
		_alpha = &alpha[0]
	}
	var _beta *float64
	if len(beta) > 0 {
		_beta = &beta[0]
	}
	var _u *complex128
	if len(u) > 0 {
		_u = &u[0]
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _ncycle *int32
	if len(ncycle) > 0 {
		_ncycle = &ncycle[0]
	}
	return isZero(C.LAPACKE_ztgsja_work((C.int)(rowMajor), (C.char)(jobu), (C.char)(jobv), (C.char)(jobq), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (C.double)(tola), (C.double)(tolb), (*C.double)(_alpha), (*C.double)(_beta), (*C.lapack_complex_double)(_u), (C.lapack_int)(ldu), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_work), (*C.lapack_int)(_ncycle)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stgsna.f.
func Stgsna(job lapack.Job, howmny byte, sel []int32, n int, a []float32, lda int, b []float32, ldb int, vl []float32, ldvl int, vr []float32, ldvr int, s, dif []float32, mm int, m []int32, work []float32, lwork int, iwork []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _dif *float32
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stgsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_s), (*C.float)(_dif), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtgsna.f.
func Dtgsna(job lapack.Job, howmny byte, sel []int32, n int, a []float64, lda int, b []float64, ldb int, vl []float64, ldvl int, vr []float64, ldvr int, s, dif []float64, mm int, m []int32, work []float64, lwork int, iwork []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _dif *float64
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtgsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_s), (*C.double)(_dif), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctgsna.f.
func Ctgsna(job lapack.Job, howmny byte, sel []int32, n int, a []complex64, lda int, b []complex64, ldb int, vl []complex64, ldvl int, vr []complex64, ldvr int, s, dif []float32, mm int, m []int32, work []complex64, lwork int, iwork []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _dif *float32
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ctgsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_s), (*C.float)(_dif), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztgsna.f.
func Ztgsna(job lapack.Job, howmny byte, sel []int32, n int, a []complex128, lda int, b []complex128, ldb int, vl []complex128, ldvl int, vr []complex128, ldvr int, s, dif []float64, mm int, m []int32, work []complex128, lwork int, iwork []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _dif *float64
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ztgsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_s), (*C.double)(_dif), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stgsyl.f.
func Stgsyl(trans blas.Transpose, ijob lapack.Job, m, n int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, d []float32, ldd int, e []float32, lde int, f []float32, ldf int, scale, dif, work []float32, lwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *float32
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float32
	if len(e) > 0 {
		_e = &e[0]
	}
	var _f *float32
	if len(f) > 0 {
		_f = &f[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _dif *float32
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stgsyl_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(ijob), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_d), (C.lapack_int)(ldd), (*C.float)(_e), (C.lapack_int)(lde), (*C.float)(_f), (C.lapack_int)(ldf), (*C.float)(_scale), (*C.float)(_dif), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtgsyl.f.
func Dtgsyl(trans blas.Transpose, ijob lapack.Job, m, n int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, d []float64, ldd int, e []float64, lde int, f []float64, ldf int, scale, dif, work []float64, lwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *float64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *float64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _f *float64
	if len(f) > 0 {
		_f = &f[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _dif *float64
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtgsyl_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(ijob), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_d), (C.lapack_int)(ldd), (*C.double)(_e), (C.lapack_int)(lde), (*C.double)(_f), (C.lapack_int)(ldf), (*C.double)(_scale), (*C.double)(_dif), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctgsyl.f.
func Ctgsyl(trans blas.Transpose, ijob lapack.Job, m, n int, a []complex64, lda int, b []complex64, ldb int, c []complex64, ldc int, d []complex64, ldd int, e []complex64, lde int, f []complex64, ldf int, scale, dif []float32, work []complex64, lwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *complex64
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex64
	if len(e) > 0 {
		_e = &e[0]
	}
	var _f *complex64
	if len(f) > 0 {
		_f = &f[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _dif *float32
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ctgsyl_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(ijob), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_d), (C.lapack_int)(ldd), (*C.lapack_complex_float)(_e), (C.lapack_int)(lde), (*C.lapack_complex_float)(_f), (C.lapack_int)(ldf), (*C.float)(_scale), (*C.float)(_dif), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztgsyl.f.
func Ztgsyl(trans blas.Transpose, ijob lapack.Job, m, n int, a []complex128, lda int, b []complex128, ldb int, c []complex128, ldc int, d []complex128, ldd int, e []complex128, lde int, f []complex128, ldf int, scale, dif []float64, work []complex128, lwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _d *complex128
	if len(d) > 0 {
		_d = &d[0]
	}
	var _e *complex128
	if len(e) > 0 {
		_e = &e[0]
	}
	var _f *complex128
	if len(f) > 0 {
		_f = &f[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	var _dif *float64
	if len(dif) > 0 {
		_dif = &dif[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_ztgsyl_work((C.int)(rowMajor), (C.char)(trans), (C.lapack_int)(ijob), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_d), (C.lapack_int)(ldd), (*C.lapack_complex_double)(_e), (C.lapack_int)(lde), (*C.lapack_complex_double)(_f), (C.lapack_int)(ldf), (*C.double)(_scale), (*C.double)(_dif), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stpcon.f.
func Stpcon(norm byte, ul blas.Uplo, d blas.Diag, n int, ap, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stpcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtpcon.f.
func Dtpcon(norm byte, ul blas.Uplo, d blas.Diag, n int, ap, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtpcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctpcon.f.
func Ctpcon(norm byte, ul blas.Uplo, d blas.Diag, n int, ap []complex64, rcond []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctpcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztpcon.f.
func Ztpcon(norm byte, ul blas.Uplo, d blas.Diag, n int, ap []complex128, rcond []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztpcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stprfs.f.
func Stprfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_stprfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtprfs.f.
func Dtprfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtprfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctprfs.f.
func Ctprfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctprfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztprfs.f.
func Ztprfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztprfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stptri.f.
func Stptri(ul blas.Uplo, d blas.Diag, n int, ap []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_stptri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtptri.f.
func Dtptri(ul blas.Uplo, d blas.Diag, n int, ap []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_dtptri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctptri.f.
func Ctptri(ul blas.Uplo, d blas.Diag, n int, ap []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_ctptri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztptri.f.
func Ztptri(ul blas.Uplo, d blas.Diag, n int, ap []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_ztptri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stptrs.f.
func Stptrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_stptrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_ap), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtptrs.f.
func Dtptrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dtptrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_ap), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctptrs.f.
func Ctptrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ctptrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztptrs.f.
func Ztptrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, ap, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ztptrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stpttf.f.
func Stpttf(transr blas.Transpose, ul blas.Uplo, n int, ap, arf []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _arf *float32
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_stpttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtpttf.f.
func Dtpttf(transr blas.Transpose, ul blas.Uplo, n int, ap, arf []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _arf *float64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_dtpttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctpttf.f.
func Ctpttf(transr blas.Transpose, ul blas.Uplo, n int, ap, arf []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _arf *complex64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_ctpttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztpttf.f.
func Ztpttf(transr blas.Transpose, ul blas.Uplo, n int, ap, arf []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _arf *complex128
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_ztpttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stpttr.f.
func Stpttr(ul blas.Uplo, n int, ap, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_stpttr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_ap), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtpttr.f.
func Dtpttr(ul blas.Uplo, n int, ap, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dtpttr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_ap), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctpttr.f.
func Ctpttr(ul blas.Uplo, n int, ap, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ctpttr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztpttr.f.
func Ztpttr(ul blas.Uplo, n int, ap, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ztpttr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strcon.f.
func Strcon(norm byte, ul blas.Uplo, d blas.Diag, n int, a []float32, lda int, rcond, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_strcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_rcond), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrcon.f.
func Dtrcon(norm byte, ul blas.Uplo, d blas.Diag, n int, a []float64, lda int, rcond, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtrcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_rcond), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrcon.f.
func Ctrcon(norm byte, ul blas.Uplo, d blas.Diag, n int, a []complex64, lda int, rcond []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float32
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctrcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.float)(_rcond), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrcon.f.
func Ztrcon(norm byte, ul blas.Uplo, d blas.Diag, n int, a []complex128, lda int, rcond []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _rcond *float64
	if len(rcond) > 0 {
		_rcond = &rcond[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztrcon_work((C.int)(rowMajor), (C.char)(norm), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.double)(_rcond), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strevc.f.
func Strevc(side blas.Side, howmny byte, sel []int32, n int, t []float32, ldt int, vl []float32, ldvl int, vr []float32, ldvr, mm int, m []int32, work []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_strevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrevc.f.
func Dtrevc(side blas.Side, howmny byte, sel []int32, n int, t []float64, ldt int, vl []float64, ldvl int, vr []float64, ldvr, mm int, m []int32, work []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtrevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrevc.f.
func Ctrevc(side blas.Side, howmny byte, sel []int32, n int, t []complex64, ldt int, vl []complex64, ldvl int, vr []complex64, ldvr, mm int, m []int32, work []complex64, rwork []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctrevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrevc.f.
func Ztrevc(side blas.Side, howmny byte, sel []int32, n int, t []complex128, ldt int, vl []complex128, ldvl int, vr []complex128, ldvr, mm int, m []int32, work []complex128, rwork []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztrevc_work((C.int)(rowMajor), (C.char)(side), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strexc.f.
func Strexc(compq lapack.Comp, n int, t []float32, ldt int, q []float32, ldq int, ifst, ilst []int32, work []float32) bool {
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _ifst *int32
	if len(ifst) > 0 {
		_ifst = &ifst[0]
	}
	var _ilst *int32
	if len(ilst) > 0 {
		_ilst = &ilst[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_strexc_work((C.int)(rowMajor), (C.char)(compq), (C.lapack_int)(n), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_q), (C.lapack_int)(ldq), (*C.lapack_int)(_ifst), (*C.lapack_int)(_ilst), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrexc.f.
func Dtrexc(compq lapack.Comp, n int, t []float64, ldt int, q []float64, ldq int, ifst, ilst []int32, work []float64) bool {
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _ifst *int32
	if len(ifst) > 0 {
		_ifst = &ifst[0]
	}
	var _ilst *int32
	if len(ilst) > 0 {
		_ilst = &ilst[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtrexc_work((C.int)(rowMajor), (C.char)(compq), (C.lapack_int)(n), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_q), (C.lapack_int)(ldq), (*C.lapack_int)(_ifst), (*C.lapack_int)(_ilst), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrexc.f.
func Ctrexc(compq lapack.Comp, n int, t []complex64, ldt int, q []complex64, ldq, ifst, ilst int) bool {
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	return isZero(C.LAPACKE_ctrexc_work((C.int)(rowMajor), (C.char)(compq), (C.lapack_int)(n), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (C.lapack_int)(ifst), (C.lapack_int)(ilst)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrexc.f.
func Ztrexc(compq lapack.Comp, n int, t []complex128, ldt int, q []complex128, ldq, ifst, ilst int) bool {
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	return isZero(C.LAPACKE_ztrexc_work((C.int)(rowMajor), (C.char)(compq), (C.lapack_int)(n), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (C.lapack_int)(ifst), (C.lapack_int)(ilst)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strrfs.f.
func Strrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []float32, lda int, b []float32, ldb int, x []float32, ldx int, ferr, berr, work []float32, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float32
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_strrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.float)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrrfs.f.
func Dtrrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []float64, lda int, b []float64, ldb int, x []float64, ldx int, ferr, berr, work []float64, iwork []int32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *float64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtrrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.double)(_work), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrrfs.f.
func Ctrrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []complex64, lda int, b []complex64, ldb int, x []complex64, ldx int, ferr, berr []float32, work []complex64, rwork []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float32
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float32
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctrrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_x), (C.lapack_int)(ldx), (*C.float)(_ferr), (*C.float)(_berr), (*C.lapack_complex_float)(_work), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrrfs.f.
func Ztrrfs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, x []complex128, ldx int, ferr, berr []float64, work []complex128, rwork []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _ferr *float64
	if len(ferr) > 0 {
		_ferr = &ferr[0]
	}
	var _berr *float64
	if len(berr) > 0 {
		_berr = &berr[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztrrfs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_x), (C.lapack_int)(ldx), (*C.double)(_ferr), (*C.double)(_berr), (*C.lapack_complex_double)(_work), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strsen.f.
func Strsen(job lapack.Job, compq lapack.Comp, sel []int32, n int, t []float32, ldt int, q []float32, ldq int, wr, wi []float32, m []int32, s, sep, work []float32, lwork int, iwork []int32, liwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *float32
	if len(q) > 0 {
		_q = &q[0]
	}
	var _wr *float32
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float32
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float32
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_strsen_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_q), (C.lapack_int)(ldq), (*C.float)(_wr), (*C.float)(_wi), (*C.lapack_int)(_m), (*C.float)(_s), (*C.float)(_sep), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrsen.f.
func Dtrsen(job lapack.Job, compq lapack.Comp, sel []int32, n int, t []float64, ldt int, q []float64, ldq int, wr, wi []float64, m []int32, s, sep, work []float64, lwork int, iwork []int32, liwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *float64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _wr *float64
	if len(wr) > 0 {
		_wr = &wr[0]
	}
	var _wi *float64
	if len(wi) > 0 {
		_wi = &wi[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float64
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtrsen_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_q), (C.lapack_int)(ldq), (*C.double)(_wr), (*C.double)(_wi), (*C.lapack_int)(_m), (*C.double)(_s), (*C.double)(_sep), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork), (C.lapack_int)(liwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrsen.f.
func Ctrsen(job lapack.Job, compq lapack.Comp, sel []int32, n int, t []complex64, ldt int, q []complex64, ldq int, w []complex64, m []int32, s, sep []float32, work []complex64, lwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _w *complex64
	if len(w) > 0 {
		_w = &w[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float32
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ctrsen_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_w), (*C.lapack_int)(_m), (*C.float)(_s), (*C.float)(_sep), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrsen.f.
func Ztrsen(job lapack.Job, compq lapack.Comp, sel []int32, n int, t []complex128, ldt int, q []complex128, ldq int, w []complex128, m []int32, s, sep []float64, work []complex128, lwork int) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _w *complex128
	if len(w) > 0 {
		_w = &w[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float64
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ztrsen_work((C.int)(rowMajor), (C.char)(job), (C.char)(compq), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_w), (*C.lapack_int)(_m), (*C.double)(_s), (*C.double)(_sep), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strsna.f.
func Strsna(job lapack.Job, howmny byte, sel []int32, n int, t []float32, ldt int, vl []float32, ldvl int, vr []float32, ldvr int, s, sep []float32, mm int, m []int32, work []float32, ldwork int, iwork []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *float32
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float32
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float32
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_strsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_vl), (C.lapack_int)(ldvl), (*C.float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_s), (*C.float)(_sep), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.float)(_work), (C.lapack_int)(ldwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrsna.f.
func Dtrsna(job lapack.Job, howmny byte, sel []int32, n int, t []float64, ldt int, vl []float64, ldvl int, vr []float64, ldvr int, s, sep []float64, mm int, m []int32, work []float64, ldwork int, iwork []int32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *float64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *float64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float64
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dtrsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_vl), (C.lapack_int)(ldvl), (*C.double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_s), (*C.double)(_sep), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.double)(_work), (C.lapack_int)(ldwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrsna.f.
func Ctrsna(job lapack.Job, howmny byte, sel []int32, n int, t []complex64, ldt int, vl []complex64, ldvl int, vr []complex64, ldvr int, s, sep []float32, mm int, m []int32, work []complex64, ldwork int, rwork []float32) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *complex64
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex64
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float32
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float32
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ctrsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_float)(_vr), (C.lapack_int)(ldvr), (*C.float)(_s), (*C.float)(_sep), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_float)(_work), (C.lapack_int)(ldwork), (*C.float)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrsna.f.
func Ztrsna(job lapack.Job, howmny byte, sel []int32, n int, t []complex128, ldt int, vl []complex128, ldvl int, vr []complex128, ldvr int, s, sep []float64, mm int, m []int32, work []complex128, ldwork int, rwork []float64) bool {
	var _sel *int32
	if len(sel) > 0 {
		_sel = &sel[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _vl *complex128
	if len(vl) > 0 {
		_vl = &vl[0]
	}
	var _vr *complex128
	if len(vr) > 0 {
		_vr = &vr[0]
	}
	var _s *float64
	if len(s) > 0 {
		_s = &s[0]
	}
	var _sep *float64
	if len(sep) > 0 {
		_sep = &sep[0]
	}
	var _m *int32
	if len(m) > 0 {
		_m = &m[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_ztrsna_work((C.int)(rowMajor), (C.char)(job), (C.char)(howmny), (*C.lapack_int)(_sel), (C.lapack_int)(n), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_vl), (C.lapack_int)(ldvl), (*C.lapack_complex_double)(_vr), (C.lapack_int)(ldvr), (*C.double)(_s), (*C.double)(_sep), (C.lapack_int)(mm), (*C.lapack_int)(_m), (*C.lapack_complex_double)(_work), (C.lapack_int)(ldwork), (*C.double)(_rwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strsyl.f.
func Strsyl(trana, tranb blas.Transpose, isgn, m, n int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, scale []float32) bool {
	switch trana {
	case blas.NoTrans:
		trana = 'N'
	case blas.Trans:
		trana = 'T'
	case blas.ConjTrans:
		trana = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch tranb {
	case blas.NoTrans:
		tranb = 'N'
	case blas.Trans:
		tranb = 'T'
	case blas.ConjTrans:
		tranb = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_strsyl_work((C.int)(rowMajor), (C.char)(trana), (C.char)(tranb), (C.lapack_int)(isgn), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrsyl.f.
func Dtrsyl(trana, tranb blas.Transpose, isgn, m, n int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, scale []float64) bool {
	switch trana {
	case blas.NoTrans:
		trana = 'N'
	case blas.Trans:
		trana = 'T'
	case blas.ConjTrans:
		trana = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch tranb {
	case blas.NoTrans:
		tranb = 'N'
	case blas.Trans:
		tranb = 'T'
	case blas.ConjTrans:
		tranb = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_dtrsyl_work((C.int)(rowMajor), (C.char)(trana), (C.char)(tranb), (C.lapack_int)(isgn), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrsyl.f.
func Ctrsyl(trana, tranb blas.Transpose, isgn, m, n int, a []complex64, lda int, b []complex64, ldb int, c []complex64, ldc int, scale []float32) bool {
	switch trana {
	case blas.NoTrans:
		trana = 'N'
	case blas.Trans:
		trana = 'T'
	case blas.ConjTrans:
		trana = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch tranb {
	case blas.NoTrans:
		tranb = 'N'
	case blas.Trans:
		tranb = 'T'
	case blas.ConjTrans:
		tranb = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _scale *float32
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_ctrsyl_work((C.int)(rowMajor), (C.char)(trana), (C.char)(tranb), (C.lapack_int)(isgn), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.float)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrsyl.f.
func Ztrsyl(trana, tranb blas.Transpose, isgn, m, n int, a []complex128, lda int, b []complex128, ldb int, c []complex128, ldc int, scale []float64) bool {
	switch trana {
	case blas.NoTrans:
		trana = 'N'
	case blas.Trans:
		trana = 'T'
	case blas.ConjTrans:
		trana = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch tranb {
	case blas.NoTrans:
		tranb = 'N'
	case blas.Trans:
		tranb = 'T'
	case blas.ConjTrans:
		tranb = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _scale *float64
	if len(scale) > 0 {
		_scale = &scale[0]
	}
	return isZero(C.LAPACKE_ztrsyl_work((C.int)(rowMajor), (C.char)(trana), (C.char)(tranb), (C.lapack_int)(isgn), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.double)(_scale)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strtri.f.
func Strtri(ul blas.Uplo, d blas.Diag, n int, a []float32, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_strtri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrtri.f.
func Dtrtri(ul blas.Uplo, d blas.Diag, n int, a []float64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dtrtri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrtri.f.
func Ctrtri(ul blas.Uplo, d blas.Diag, n int, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ctrtri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrtri.f.
func Ztrtri(ul blas.Uplo, d blas.Diag, n int, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ztrtri_work((C.int)(rowMajor), (C.char)(ul), (C.char)(d), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strtrs.f.
func Strtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []float32, lda int, b []float32, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_strtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrtrs.f.
func Dtrtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []float64, lda int, b []float64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_dtrtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrtrs.f.
func Ctrtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []complex64, lda int, b []complex64, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ctrtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrtrs.f.
func Ztrtrs(ul blas.Uplo, trans blas.Transpose, d blas.Diag, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch d {
	case blas.Unit:
		d = 'U'
	case blas.NonUnit:
		d = 'N'
	default:
		panic("lapack: illegal diagonal")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	return isZero(C.LAPACKE_ztrtrs_work((C.int)(rowMajor), (C.char)(ul), (C.char)(trans), (C.char)(d), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strttf.f.
func Strttf(transr blas.Transpose, ul blas.Uplo, n int, a []float32, lda int, arf []float32) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _arf *float32
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_strttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrttf.f.
func Dtrttf(transr blas.Transpose, ul blas.Uplo, n int, a []float64, lda int, arf []float64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _arf *float64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_dtrttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrttf.f.
func Ctrttf(transr blas.Transpose, ul blas.Uplo, n int, a []complex64, lda int, arf []complex64) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _arf *complex64
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_ctrttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrttf.f.
func Ztrttf(transr blas.Transpose, ul blas.Uplo, n int, a []complex128, lda int, arf []complex128) bool {
	switch transr {
	case blas.NoTrans:
		transr = 'N'
	case blas.Trans:
		transr = 'T'
	case blas.ConjTrans:
		transr = 'C'
	default:
		panic("lapack: bad trans")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _arf *complex128
	if len(arf) > 0 {
		_arf = &arf[0]
	}
	return isZero(C.LAPACKE_ztrttf_work((C.int)(rowMajor), (C.char)(transr), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_arf)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/strttp.f.
func Strttp(ul blas.Uplo, n int, a []float32, lda int, ap []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ap *float32
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_strttp_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtrttp.f.
func Dtrttp(ul blas.Uplo, n int, a []float64, lda int, ap []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ap *float64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_dtrttp_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctrttp.f.
func Ctrttp(ul blas.Uplo, n int, a []complex64, lda int, ap []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_ctrttp_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztrttp.f.
func Ztrttp(ul blas.Uplo, n int, a []complex128, lda int, ap []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	return isZero(C.LAPACKE_ztrttp_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_ap)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stzrzf.f.
func Stzrzf(m, n int, a []float32, lda int, tau, work []float32, lwork int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float32
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_stzrzf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_tau), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtzrzf.f.
func Dtzrzf(m, n int, a []float64, lda int, tau, work []float64, lwork int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *float64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtzrzf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_tau), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctzrzf.f.
func Ctzrzf(m, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ctzrzf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztzrzf.f.
func Ztzrzf(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ztzrzf_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cungbr.f.
func Cungbr(vect byte, m, n, k int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cungbr_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zungbr.f.
func Zungbr(vect byte, m, n, k int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zungbr_work((C.int)(rowMajor), (C.char)(vect), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunghr.f.
func Cunghr(n, ilo, ihi int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunghr_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunghr.f.
func Zunghr(n, ilo, ihi int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunghr_work((C.int)(rowMajor), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunglq.f.
func Cunglq(m, n, k int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunglq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunglq.f.
func Zunglq(m, n, k int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunglq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cungql.f.
func Cungql(m, n, k int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cungql_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zungql.f.
func Zungql(m, n, k int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zungql_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cungqr.f.
func Cungqr(m, n, k int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cungqr_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zungqr.f.
func Zungqr(m, n, k int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zungqr_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cungrq.f.
func Cungrq(m, n, k int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cungrq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zungrq.f.
func Zungrq(m, n, k int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zungrq_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cungtr.f.
func Cungtr(ul blas.Uplo, n int, a []complex64, lda int, tau, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cungtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zungtr.f.
func Zungtr(ul blas.Uplo, n int, a []complex128, lda int, tau, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zungtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmbr.f.
func Cunmbr(vect byte, side blas.Side, trans blas.Transpose, m, n, k int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmbr_work((C.int)(rowMajor), (C.char)(vect), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmbr.f.
func Zunmbr(vect byte, side blas.Side, trans blas.Transpose, m, n, k int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmbr_work((C.int)(rowMajor), (C.char)(vect), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmhr.f.
func Cunmhr(side blas.Side, trans blas.Transpose, m, n, ilo, ihi int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmhr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmhr.f.
func Zunmhr(side blas.Side, trans blas.Transpose, m, n, ilo, ihi int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmhr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(ilo), (C.lapack_int)(ihi), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmlq.f.
func Cunmlq(side blas.Side, trans blas.Transpose, m, n, k int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmlq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmlq.f.
func Zunmlq(side blas.Side, trans blas.Transpose, m, n, k int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmlq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmql.f.
func Cunmql(side blas.Side, trans blas.Transpose, m, n, k int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmql_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmql.f.
func Zunmql(side blas.Side, trans blas.Transpose, m, n, k int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmql_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmqr.f.
func Cunmqr(side blas.Side, trans blas.Transpose, m, n, k int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmqr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmqr.f.
func Zunmqr(side blas.Side, trans blas.Transpose, m, n, k int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmqr_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmrq.f.
func Cunmrq(side blas.Side, trans blas.Transpose, m, n, k int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmrq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmrq.f.
func Zunmrq(side blas.Side, trans blas.Transpose, m, n, k int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmrq_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmrz.f.
func Cunmrz(side blas.Side, trans blas.Transpose, m, n, k, l int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmrz_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmrz.f.
func Zunmrz(side blas.Side, trans blas.Transpose, m, n, k, l int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmrz_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunmtr.f.
func Cunmtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, a []complex64, lda int, tau, c []complex64, ldc int, work []complex64, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunmtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunmtr.f.
func Zunmtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, a []complex128, lda int, tau, c []complex128, ldc int, work []complex128, lwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunmtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cupgtr.f.
func Cupgtr(ul blas.Uplo, n int, ap, tau, q []complex64, ldq int, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _q *complex64
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cupgtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zupgtr.f.
func Zupgtr(ul blas.Uplo, n int, ap, tau, q []complex128, ldq int, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _q *complex128
	if len(q) > 0 {
		_q = &q[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zupgtr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_q), (C.lapack_int)(ldq), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cupmtr.f.
func Cupmtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, ap, tau, c []complex64, ldc int, work []complex64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ap *complex64
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *complex64
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cupmtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_ap), (*C.lapack_complex_float)(_tau), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zupmtr.f.
func Zupmtr(side blas.Side, ul blas.Uplo, trans blas.Transpose, m, n int, ap, tau, c []complex128, ldc int, work []complex128) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _ap *complex128
	if len(ap) > 0 {
		_ap = &ap[0]
	}
	var _tau *complex128
	if len(tau) > 0 {
		_tau = &tau[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zupmtr_work((C.int)(rowMajor), (C.char)(side), (C.char)(ul), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_ap), (*C.lapack_complex_double)(_tau), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cbbcsd.f.
func Cbbcsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, m, p, q int, theta, phi []float32, u1 []complex64, ldu1 int, u2 []complex64, ldu2 int, v1t []complex64, ldv1t int, v2t []complex64, ldv2t int, b11d, b11e, b12d, b12e, b21d, b21e, b22d, b22e, rwork []float32, lrwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float32
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _u1 *complex64
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *complex64
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *complex64
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *complex64
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _b11d *float32
	if len(b11d) > 0 {
		_b11d = &b11d[0]
	}
	var _b11e *float32
	if len(b11e) > 0 {
		_b11e = &b11e[0]
	}
	var _b12d *float32
	if len(b12d) > 0 {
		_b12d = &b12d[0]
	}
	var _b12e *float32
	if len(b12e) > 0 {
		_b12e = &b12e[0]
	}
	var _b21d *float32
	if len(b21d) > 0 {
		_b21d = &b21d[0]
	}
	var _b21e *float32
	if len(b21e) > 0 {
		_b21e = &b21e[0]
	}
	var _b22d *float32
	if len(b22d) > 0 {
		_b22d = &b22d[0]
	}
	var _b22e *float32
	if len(b22e) > 0 {
		_b22e = &b22e[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_cbbcsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.float)(_theta), (*C.float)(_phi), (*C.lapack_complex_float)(_u1), (C.lapack_int)(ldu1), (*C.lapack_complex_float)(_u2), (C.lapack_int)(ldu2), (*C.lapack_complex_float)(_v1t), (C.lapack_int)(ldv1t), (*C.lapack_complex_float)(_v2t), (C.lapack_int)(ldv2t), (*C.float)(_b11d), (*C.float)(_b11e), (*C.float)(_b12d), (*C.float)(_b12e), (*C.float)(_b21d), (*C.float)(_b21e), (*C.float)(_b22d), (*C.float)(_b22e), (*C.float)(_rwork), (C.lapack_int)(lrwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cheswapr.f.
func Cheswapr(ul blas.Uplo, n int, a []complex64, i1, i2 int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_cheswapr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(i1), (C.lapack_int)(i2)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetri2.f.
func Chetri2(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chetri2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetri2x.f.
func Chetri2x(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64, nb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chetri2x_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(nb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/chetrs2.f.
func Chetrs2(ul blas.Uplo, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_chetrs2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csyconv.f.
func Csyconv(ul blas.Uplo, way byte, n int, a []complex64, lda int, ipiv []int32, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csyconv_work((C.int)(rowMajor), (C.char)(ul), (C.char)(way), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csyswapr.f.
func Csyswapr(ul blas.Uplo, n int, a []complex64, i1, i2 int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_csyswapr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(i1), (C.lapack_int)(i2)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csytri2.f.
func Csytri2(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csytri2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csytri2x.f.
func Csytri2x(ul blas.Uplo, n int, a []complex64, lda int, ipiv []int32, work []complex64, nb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csytri2x_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(nb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csytrs2.f.
func Csytrs2(ul blas.Uplo, n, nrhs int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int, work []complex64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_csytrs2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cunbdb.f.
func Cunbdb(trans blas.Transpose, signs byte, m, p, q int, x11 []complex64, ldx11 int, x12 []complex64, ldx12 int, x21 []complex64, ldx21 int, x22 []complex64, ldx22 int, theta, phi []float32, taup1, taup2, tauq1, tauq2, work []complex64, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *complex64
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *complex64
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *complex64
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *complex64
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float32
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _taup1 *complex64
	if len(taup1) > 0 {
		_taup1 = &taup1[0]
	}
	var _taup2 *complex64
	if len(taup2) > 0 {
		_taup2 = &taup2[0]
	}
	var _tauq1 *complex64
	if len(tauq1) > 0 {
		_tauq1 = &tauq1[0]
	}
	var _tauq2 *complex64
	if len(tauq2) > 0 {
		_tauq2 = &tauq2[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cunbdb_work((C.int)(rowMajor), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.lapack_complex_float)(_x11), (C.lapack_int)(ldx11), (*C.lapack_complex_float)(_x12), (C.lapack_int)(ldx12), (*C.lapack_complex_float)(_x21), (C.lapack_int)(ldx21), (*C.lapack_complex_float)(_x22), (C.lapack_int)(ldx22), (*C.float)(_theta), (*C.float)(_phi), (*C.lapack_complex_float)(_taup1), (*C.lapack_complex_float)(_taup2), (*C.lapack_complex_float)(_tauq1), (*C.lapack_complex_float)(_tauq2), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cuncsd.f.
func Cuncsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, signs byte, m, p, q int, x11 []complex64, ldx11 int, x12 []complex64, ldx12 int, x21 []complex64, ldx21 int, x22 []complex64, ldx22 int, theta []float32, u1 []complex64, ldu1 int, u2 []complex64, ldu2 int, v1t []complex64, ldv1t int, v2t []complex64, ldv2t int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *complex64
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *complex64
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *complex64
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *complex64
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *complex64
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *complex64
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *complex64
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *complex64
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cuncsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.lapack_complex_float)(_x11), (C.lapack_int)(ldx11), (*C.lapack_complex_float)(_x12), (C.lapack_int)(ldx12), (*C.lapack_complex_float)(_x21), (C.lapack_int)(ldx21), (*C.lapack_complex_float)(_x22), (C.lapack_int)(ldx22), (*C.float)(_theta), (*C.lapack_complex_float)(_u1), (C.lapack_int)(ldu1), (*C.lapack_complex_float)(_u2), (C.lapack_int)(ldu2), (*C.lapack_complex_float)(_v1t), (C.lapack_int)(ldv1t), (*C.lapack_complex_float)(_v2t), (C.lapack_int)(ldv2t), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cuncsd2by1.f.
func Cuncsd2by1(jobu1, jobu2, jobv1t lapack.Job, m, p, q int, x11 []complex64, ldx11 int, x21 []complex64, ldx21 int, theta, u1 []complex64, ldu1 int, u2 []complex64, ldu2 int, v1t []complex64, ldv1t int, work []complex64, lwork int, rwork []float32, lrwork int, iwork []int32) bool {
	var _x11 *complex64
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x21 *complex64
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _theta *complex64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *complex64
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *complex64
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *complex64
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float32
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_cuncsd2by1_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.lapack_complex_float)(_x11), (C.lapack_int)(ldx11), (*C.lapack_complex_float)(_x21), (C.lapack_int)(ldx21), (*C.lapack_complex_float)(_theta), (*C.lapack_complex_float)(_u1), (C.lapack_int)(ldu1), (*C.lapack_complex_float)(_u2), (C.lapack_int)(ldu2), (*C.lapack_complex_float)(_v1t), (C.lapack_int)(ldv1t), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork), (*C.float)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dbbcsd.f.
func Dbbcsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, m, p, q int, theta, phi, u1 []float64, ldu1 int, u2 []float64, ldu2 int, v1t []float64, ldv1t int, v2t []float64, ldv2t int, b11d, b11e, b12d, b12e, b21d, b21e, b22d, b22e, work []float64, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float64
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _u1 *float64
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *float64
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *float64
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *float64
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _b11d *float64
	if len(b11d) > 0 {
		_b11d = &b11d[0]
	}
	var _b11e *float64
	if len(b11e) > 0 {
		_b11e = &b11e[0]
	}
	var _b12d *float64
	if len(b12d) > 0 {
		_b12d = &b12d[0]
	}
	var _b12e *float64
	if len(b12e) > 0 {
		_b12e = &b12e[0]
	}
	var _b21d *float64
	if len(b21d) > 0 {
		_b21d = &b21d[0]
	}
	var _b21e *float64
	if len(b21e) > 0 {
		_b21e = &b21e[0]
	}
	var _b22d *float64
	if len(b22d) > 0 {
		_b22d = &b22d[0]
	}
	var _b22e *float64
	if len(b22e) > 0 {
		_b22e = &b22e[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dbbcsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.double)(_theta), (*C.double)(_phi), (*C.double)(_u1), (C.lapack_int)(ldu1), (*C.double)(_u2), (C.lapack_int)(ldu2), (*C.double)(_v1t), (C.lapack_int)(ldv1t), (*C.double)(_v2t), (C.lapack_int)(ldv2t), (*C.double)(_b11d), (*C.double)(_b11e), (*C.double)(_b12d), (*C.double)(_b12e), (*C.double)(_b21d), (*C.double)(_b21e), (*C.double)(_b22d), (*C.double)(_b22e), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorbdb.f.
func Dorbdb(trans blas.Transpose, signs byte, m, p, q int, x11 []float64, ldx11 int, x12 []float64, ldx12 int, x21 []float64, ldx21 int, x22 []float64, ldx22 int, theta, phi, taup1, taup2, tauq1, tauq2, work []float64, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *float64
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *float64
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *float64
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *float64
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float64
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _taup1 *float64
	if len(taup1) > 0 {
		_taup1 = &taup1[0]
	}
	var _taup2 *float64
	if len(taup2) > 0 {
		_taup2 = &taup2[0]
	}
	var _tauq1 *float64
	if len(tauq1) > 0 {
		_tauq1 = &tauq1[0]
	}
	var _tauq2 *float64
	if len(tauq2) > 0 {
		_tauq2 = &tauq2[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dorbdb_work((C.int)(rowMajor), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.double)(_x11), (C.lapack_int)(ldx11), (*C.double)(_x12), (C.lapack_int)(ldx12), (*C.double)(_x21), (C.lapack_int)(ldx21), (*C.double)(_x22), (C.lapack_int)(ldx22), (*C.double)(_theta), (*C.double)(_phi), (*C.double)(_taup1), (*C.double)(_taup2), (*C.double)(_tauq1), (*C.double)(_tauq2), (*C.double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorcsd.f.
func Dorcsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, signs byte, m, p, q int, x11 []float64, ldx11 int, x12 []float64, ldx12 int, x21 []float64, ldx21 int, x22 []float64, ldx22 int, theta, u1 []float64, ldu1 int, u2 []float64, ldu2 int, v1t []float64, ldv1t int, v2t []float64, ldv2t int, work []float64, lwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *float64
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *float64
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *float64
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *float64
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *float64
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *float64
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *float64
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *float64
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dorcsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.double)(_x11), (C.lapack_int)(ldx11), (*C.double)(_x12), (C.lapack_int)(ldx12), (*C.double)(_x21), (C.lapack_int)(ldx21), (*C.double)(_x22), (C.lapack_int)(ldx22), (*C.double)(_theta), (*C.double)(_u1), (C.lapack_int)(ldu1), (*C.double)(_u2), (C.lapack_int)(ldu2), (*C.double)(_v1t), (C.lapack_int)(ldv1t), (*C.double)(_v2t), (C.lapack_int)(ldv2t), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dorcsd2by1.f.
func Dorcsd2by1(jobu1, jobu2, jobv1t lapack.Job, m, p, q int, x11 []float64, ldx11 int, x21 []float64, ldx21 int, theta, u1 []float64, ldu1 int, u2 []float64, ldu2 int, v1t []float64, ldv1t int, work []float64, lwork int, iwork []int32) bool {
	var _x11 *float64
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x21 *float64
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *float64
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *float64
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *float64
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_dorcsd2by1_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.double)(_x11), (C.lapack_int)(ldx11), (*C.double)(_x21), (C.lapack_int)(ldx21), (*C.double)(_theta), (*C.double)(_u1), (C.lapack_int)(ldu1), (*C.double)(_u2), (C.lapack_int)(ldu2), (*C.double)(_v1t), (C.lapack_int)(ldv1t), (*C.double)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyconv.f.
func Dsyconv(ul blas.Uplo, way byte, n int, a []float64, lda int, ipiv []int32, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsyconv_work((C.int)(rowMajor), (C.char)(ul), (C.char)(way), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsyswapr.f.
func Dsyswapr(ul blas.Uplo, n int, a []float64, i1, i2 int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_dsyswapr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(i1), (C.lapack_int)(i2)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytri2.f.
func Dsytri2(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsytri2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytri2x.f.
func Dsytri2x(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32, work []float64, nb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsytri2x_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_work), (C.lapack_int)(nb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dsytrs2.f.
func Dsytrs2(ul blas.Uplo, n, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int, work []float64) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dsytrs2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sbbcsd.f.
func Sbbcsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, m, p, q int, theta, phi, u1 []float32, ldu1 int, u2 []float32, ldu2 int, v1t []float32, ldv1t int, v2t []float32, ldv2t int, b11d, b11e, b12d, b12e, b21d, b21e, b22d, b22e, work []float32, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float32
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _u1 *float32
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *float32
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *float32
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *float32
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _b11d *float32
	if len(b11d) > 0 {
		_b11d = &b11d[0]
	}
	var _b11e *float32
	if len(b11e) > 0 {
		_b11e = &b11e[0]
	}
	var _b12d *float32
	if len(b12d) > 0 {
		_b12d = &b12d[0]
	}
	var _b12e *float32
	if len(b12e) > 0 {
		_b12e = &b12e[0]
	}
	var _b21d *float32
	if len(b21d) > 0 {
		_b21d = &b21d[0]
	}
	var _b21e *float32
	if len(b21e) > 0 {
		_b21e = &b21e[0]
	}
	var _b22d *float32
	if len(b22d) > 0 {
		_b22d = &b22d[0]
	}
	var _b22e *float32
	if len(b22e) > 0 {
		_b22e = &b22e[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sbbcsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.float)(_theta), (*C.float)(_phi), (*C.float)(_u1), (C.lapack_int)(ldu1), (*C.float)(_u2), (C.lapack_int)(ldu2), (*C.float)(_v1t), (C.lapack_int)(ldv1t), (*C.float)(_v2t), (C.lapack_int)(ldv2t), (*C.float)(_b11d), (*C.float)(_b11e), (*C.float)(_b12d), (*C.float)(_b12e), (*C.float)(_b21d), (*C.float)(_b21e), (*C.float)(_b22d), (*C.float)(_b22e), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorbdb.f.
func Sorbdb(trans blas.Transpose, signs byte, m, p, q int, x11 []float32, ldx11 int, x12 []float32, ldx12 int, x21 []float32, ldx21 int, x22 []float32, ldx22 int, theta, phi, taup1, taup2, tauq1, tauq2, work []float32, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *float32
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *float32
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *float32
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *float32
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float32
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _taup1 *float32
	if len(taup1) > 0 {
		_taup1 = &taup1[0]
	}
	var _taup2 *float32
	if len(taup2) > 0 {
		_taup2 = &taup2[0]
	}
	var _tauq1 *float32
	if len(tauq1) > 0 {
		_tauq1 = &tauq1[0]
	}
	var _tauq2 *float32
	if len(tauq2) > 0 {
		_tauq2 = &tauq2[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sorbdb_work((C.int)(rowMajor), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.float)(_x11), (C.lapack_int)(ldx11), (*C.float)(_x12), (C.lapack_int)(ldx12), (*C.float)(_x21), (C.lapack_int)(ldx21), (*C.float)(_x22), (C.lapack_int)(ldx22), (*C.float)(_theta), (*C.float)(_phi), (*C.float)(_taup1), (*C.float)(_taup2), (*C.float)(_tauq1), (*C.float)(_tauq2), (*C.float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorcsd.f.
func Sorcsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, signs byte, m, p, q int, x11 []float32, ldx11 int, x12 []float32, ldx12 int, x21 []float32, ldx21 int, x22 []float32, ldx22 int, theta, u1 []float32, ldu1 int, u2 []float32, ldu2 int, v1t []float32, ldv1t int, v2t []float32, ldv2t int, work []float32, lwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *float32
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *float32
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *float32
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *float32
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *float32
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *float32
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *float32
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *float32
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sorcsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.float)(_x11), (C.lapack_int)(ldx11), (*C.float)(_x12), (C.lapack_int)(ldx12), (*C.float)(_x21), (C.lapack_int)(ldx21), (*C.float)(_x22), (C.lapack_int)(ldx22), (*C.float)(_theta), (*C.float)(_u1), (C.lapack_int)(ldu1), (*C.float)(_u2), (C.lapack_int)(ldu2), (*C.float)(_v1t), (C.lapack_int)(ldv1t), (*C.float)(_v2t), (C.lapack_int)(ldv2t), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sorcsd2by1.f.
func Sorcsd2by1(jobu1, jobu2, jobv1t lapack.Job, m, p, q int, x11 []float32, ldx11 int, x21 []float32, ldx21 int, theta, u1 []float32, ldu1 int, u2 []float32, ldu2 int, v1t []float32, ldv1t int, work []float32, lwork int, iwork []int32) bool {
	var _x11 *float32
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x21 *float32
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _theta *float32
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *float32
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *float32
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *float32
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_sorcsd2by1_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.float)(_x11), (C.lapack_int)(ldx11), (*C.float)(_x21), (C.lapack_int)(ldx21), (*C.float)(_theta), (*C.float)(_u1), (C.lapack_int)(ldu1), (*C.float)(_u2), (C.lapack_int)(ldu2), (*C.float)(_v1t), (C.lapack_int)(ldv1t), (*C.float)(_work), (C.lapack_int)(lwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyconv.f.
func Ssyconv(ul blas.Uplo, way byte, n int, a []float32, lda int, ipiv []int32, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssyconv_work((C.int)(rowMajor), (C.char)(ul), (C.char)(way), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssyswapr.f.
func Ssyswapr(ul blas.Uplo, n int, a []float32, i1, i2 int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_ssyswapr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(i1), (C.lapack_int)(i2)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytri2.f.
func Ssytri2(ul blas.Uplo, n int, a []float32, lda int, ipiv []int32, work []complex64, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssytri2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_float)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytri2x.f.
func Ssytri2x(ul blas.Uplo, n int, a []float32, lda int, ipiv []int32, work []float32, nb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssytri2x_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_work), (C.lapack_int)(nb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ssytrs2.f.
func Ssytrs2(ul blas.Uplo, n, nrhs int, a []float32, lda int, ipiv []int32, b []float32, ldb int, work []float32) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ssytrs2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.float)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zbbcsd.f.
func Zbbcsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, m, p, q int, theta, phi []float64, u1 []complex128, ldu1 int, u2 []complex128, ldu2 int, v1t []complex128, ldv1t int, v2t []complex128, ldv2t int, b11d, b11e, b12d, b12e, b21d, b21e, b22d, b22e, rwork []float64, lrwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float64
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _u1 *complex128
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *complex128
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *complex128
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *complex128
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _b11d *float64
	if len(b11d) > 0 {
		_b11d = &b11d[0]
	}
	var _b11e *float64
	if len(b11e) > 0 {
		_b11e = &b11e[0]
	}
	var _b12d *float64
	if len(b12d) > 0 {
		_b12d = &b12d[0]
	}
	var _b12e *float64
	if len(b12e) > 0 {
		_b12e = &b12e[0]
	}
	var _b21d *float64
	if len(b21d) > 0 {
		_b21d = &b21d[0]
	}
	var _b21e *float64
	if len(b21e) > 0 {
		_b21e = &b21e[0]
	}
	var _b22d *float64
	if len(b22d) > 0 {
		_b22d = &b22d[0]
	}
	var _b22e *float64
	if len(b22e) > 0 {
		_b22e = &b22e[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	return isZero(C.LAPACKE_zbbcsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.double)(_theta), (*C.double)(_phi), (*C.lapack_complex_double)(_u1), (C.lapack_int)(ldu1), (*C.lapack_complex_double)(_u2), (C.lapack_int)(ldu2), (*C.lapack_complex_double)(_v1t), (C.lapack_int)(ldv1t), (*C.lapack_complex_double)(_v2t), (C.lapack_int)(ldv2t), (*C.double)(_b11d), (*C.double)(_b11e), (*C.double)(_b12d), (*C.double)(_b12e), (*C.double)(_b21d), (*C.double)(_b21e), (*C.double)(_b22d), (*C.double)(_b22e), (*C.double)(_rwork), (C.lapack_int)(lrwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zheswapr.f.
func Zheswapr(ul blas.Uplo, n int, a []complex128, i1, i2 int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zheswapr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(i1), (C.lapack_int)(i2)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetri2.f.
func Zhetri2(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhetri2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetri2x.f.
func Zhetri2x(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128, nb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhetri2x_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(nb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zhetrs2.f.
func Zhetrs2(ul blas.Uplo, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zhetrs2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsyconv.f.
func Zsyconv(ul blas.Uplo, way byte, n int, a []complex128, lda int, ipiv []int32, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsyconv_work((C.int)(rowMajor), (C.char)(ul), (C.char)(way), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsyswapr.f.
func Zsyswapr(ul blas.Uplo, n int, a []complex128, i1, i2 int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zsyswapr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(i1), (C.lapack_int)(i2)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsytri2.f.
func Zsytri2(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128, lwork int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsytri2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsytri2x.f.
func Zsytri2x(ul blas.Uplo, n int, a []complex128, lda int, ipiv []int32, work []complex128, nb int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsytri2x_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_work), (C.lapack_int)(nb)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsytrs2.f.
func Zsytrs2(ul blas.Uplo, n, nrhs int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int, work []complex128) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _ipiv *int32
	if len(ipiv) > 0 {
		_ipiv = &ipiv[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zsytrs2_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_int)(nrhs), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_int)(_ipiv), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zunbdb.f.
func Zunbdb(trans blas.Transpose, signs byte, m, p, q int, x11 []complex128, ldx11 int, x12 []complex128, ldx12 int, x21 []complex128, ldx21 int, x22 []complex128, ldx22 int, theta, phi []float64, taup1, taup2, tauq1, tauq2, work []complex128, lwork int) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *complex128
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *complex128
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *complex128
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *complex128
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _phi *float64
	if len(phi) > 0 {
		_phi = &phi[0]
	}
	var _taup1 *complex128
	if len(taup1) > 0 {
		_taup1 = &taup1[0]
	}
	var _taup2 *complex128
	if len(taup2) > 0 {
		_taup2 = &taup2[0]
	}
	var _tauq1 *complex128
	if len(tauq1) > 0 {
		_tauq1 = &tauq1[0]
	}
	var _tauq2 *complex128
	if len(tauq2) > 0 {
		_tauq2 = &tauq2[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zunbdb_work((C.int)(rowMajor), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.lapack_complex_double)(_x11), (C.lapack_int)(ldx11), (*C.lapack_complex_double)(_x12), (C.lapack_int)(ldx12), (*C.lapack_complex_double)(_x21), (C.lapack_int)(ldx21), (*C.lapack_complex_double)(_x22), (C.lapack_int)(ldx22), (*C.double)(_theta), (*C.double)(_phi), (*C.lapack_complex_double)(_taup1), (*C.lapack_complex_double)(_taup2), (*C.lapack_complex_double)(_tauq1), (*C.lapack_complex_double)(_tauq2), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zuncsd.f.
func Zuncsd(jobu1, jobu2, jobv1t, jobv2t lapack.Job, trans blas.Transpose, signs byte, m, p, q int, x11 []complex128, ldx11 int, x12 []complex128, ldx12 int, x21 []complex128, ldx21 int, x22 []complex128, ldx22 int, theta []float64, u1 []complex128, ldu1 int, u2 []complex128, ldu2 int, v1t []complex128, ldv1t int, v2t []complex128, ldv2t int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32) bool {
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _x11 *complex128
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x12 *complex128
	if len(x12) > 0 {
		_x12 = &x12[0]
	}
	var _x21 *complex128
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _x22 *complex128
	if len(x22) > 0 {
		_x22 = &x22[0]
	}
	var _theta *float64
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *complex128
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *complex128
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *complex128
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _v2t *complex128
	if len(v2t) > 0 {
		_v2t = &v2t[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zuncsd_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.char)(jobv2t), (C.char)(trans), (C.char)(signs), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.lapack_complex_double)(_x11), (C.lapack_int)(ldx11), (*C.lapack_complex_double)(_x12), (C.lapack_int)(ldx12), (*C.lapack_complex_double)(_x21), (C.lapack_int)(ldx21), (*C.lapack_complex_double)(_x22), (C.lapack_int)(ldx22), (*C.double)(_theta), (*C.lapack_complex_double)(_u1), (C.lapack_int)(ldu1), (*C.lapack_complex_double)(_u2), (C.lapack_int)(ldu2), (*C.lapack_complex_double)(_v1t), (C.lapack_int)(ldv1t), (*C.lapack_complex_double)(_v2t), (C.lapack_int)(ldv2t), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zuncsd2by1.f.
func Zuncsd2by1(jobu1, jobu2, jobv1t lapack.Job, m, p, q int, x11 []complex128, ldx11 int, x21 []complex128, ldx21 int, theta, u1 []complex128, ldu1 int, u2 []complex128, ldu2 int, v1t []complex128, ldv1t int, work []complex128, lwork int, rwork []float64, lrwork int, iwork []int32) bool {
	var _x11 *complex128
	if len(x11) > 0 {
		_x11 = &x11[0]
	}
	var _x21 *complex128
	if len(x21) > 0 {
		_x21 = &x21[0]
	}
	var _theta *complex128
	if len(theta) > 0 {
		_theta = &theta[0]
	}
	var _u1 *complex128
	if len(u1) > 0 {
		_u1 = &u1[0]
	}
	var _u2 *complex128
	if len(u2) > 0 {
		_u2 = &u2[0]
	}
	var _v1t *complex128
	if len(v1t) > 0 {
		_v1t = &v1t[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	var _rwork *float64
	if len(rwork) > 0 {
		_rwork = &rwork[0]
	}
	var _iwork *int32
	if len(iwork) > 0 {
		_iwork = &iwork[0]
	}
	return isZero(C.LAPACKE_zuncsd2by1_work((C.int)(rowMajor), (C.char)(jobu1), (C.char)(jobu2), (C.char)(jobv1t), (C.lapack_int)(m), (C.lapack_int)(p), (C.lapack_int)(q), (*C.lapack_complex_double)(_x11), (C.lapack_int)(ldx11), (*C.lapack_complex_double)(_x21), (C.lapack_int)(ldx21), (*C.lapack_complex_double)(_theta), (*C.lapack_complex_double)(_u1), (C.lapack_int)(ldu1), (*C.lapack_complex_double)(_u2), (C.lapack_int)(ldu2), (*C.lapack_complex_double)(_v1t), (C.lapack_int)(ldv1t), (*C.lapack_complex_double)(_work), (C.lapack_int)(lwork), (*C.double)(_rwork), (C.lapack_int)(lrwork), (*C.lapack_int)(_iwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgemqrt.f.
func Sgemqrt(side blas.Side, trans blas.Transpose, m, n, k, nb int, v []float32, ldv int, t []float32, ldt int, c []float32, ldc int, work []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *float32
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgemqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(nb), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_c), (C.lapack_int)(ldc), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgemqrt.f.
func Dgemqrt(side blas.Side, trans blas.Transpose, m, n, k, nb int, v []float64, ldv int, t []float64, ldt int, c []float64, ldc int, work []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *float64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgemqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(nb), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_c), (C.lapack_int)(ldc), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgemqrt.f.
func Cgemqrt(side blas.Side, trans blas.Transpose, m, n, k, nb int, v []complex64, ldv int, t []complex64, ldt int, c []complex64, ldc int, work []complex64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *complex64
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgemqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(nb), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgemqrt.f.
func Zgemqrt(side blas.Side, trans blas.Transpose, m, n, k, nb int, v []complex128, ldv int, t []complex128, ldt int, c []complex128, ldc int, work []complex128) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _c *complex128
	if len(c) > 0 {
		_c = &c[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgemqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(nb), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_c), (C.lapack_int)(ldc), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqrt.f.
func Sgeqrt(m, n, nb int, a []float32, lda int, t []float32, ldt int, work []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_sgeqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nb), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqrt.f.
func Dgeqrt(m, n, nb int, a []float64, lda int, t []float64, ldt int, work []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dgeqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nb), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqrt.f.
func Cgeqrt(m, n, nb int, a []complex64, lda int, t []complex64, ldt int, work []complex64) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_cgeqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nb), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqrt.f.
func Zgeqrt(m, n, nb int, a []complex128, lda int, t []complex128, ldt int, work []complex128) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_zgeqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(nb), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqrt2.f.
func Sgeqrt2(m, n int, a []float32, lda int, t []float32, ldt int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_sgeqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqrt2.f.
func Dgeqrt2(m, n int, a []float64, lda int, t []float64, ldt int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_dgeqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqrt2.f.
func Cgeqrt2(m, n int, a []complex64, lda int, t []complex64, ldt int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_cgeqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqrt2.f.
func Zgeqrt2(m, n int, a []complex128, lda int, t []complex128, ldt int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_zgeqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/sgeqrt3.f.
func Sgeqrt3(m, n int, a []float32, lda int, t []float32, ldt int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_sgeqrt3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dgeqrt3.f.
func Dgeqrt3(m, n int, a []float64, lda int, t []float64, ldt int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_dgeqrt3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/cgeqrt3.f.
func Cgeqrt3(m, n int, a []complex64, lda int, t []complex64, ldt int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_cgeqrt3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zgeqrt3.f.
func Zgeqrt3(m, n int, a []complex128, lda int, t []complex128, ldt int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_zgeqrt3_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stpmqrt.f.
func Stpmqrt(side blas.Side, trans blas.Transpose, m, n, k, l, nb int, v []float32, ldv int, t []float32, ldt int, a []float32, lda int, b []float32, ldb int, work []float32) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_stpmqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtpmqrt.f.
func Dtpmqrt(side blas.Side, trans blas.Transpose, m, n, k, l, nb int, v []float64, ldv int, t []float64, ldt int, a []float64, lda int, b []float64, ldb int, work []float64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtpmqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctpmqrt.f.
func Ctpmqrt(side blas.Side, trans blas.Transpose, m, n, k, l, nb int, v []complex64, ldv int, t []complex64, ldt int, a []complex64, lda int, b []complex64, ldb int, work []complex64) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ctpmqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztpmqrt.f.
func Ztpmqrt(side blas.Side, trans blas.Transpose, m, n, k, l, nb int, v []complex128, ldv int, t []complex128, ldt int, a []complex128, lda int, b []complex128, ldb int, work []complex128) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ztpmqrt_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stpqrt.f.
func Stpqrt(m, n, l, nb int, a []float32, lda int, b []float32, ldb int, t []float32, ldt int, work []float32) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_stpqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtpqrt.f.
func Dtpqrt(m, n, l, nb int, a []float64, lda int, b []float64, ldb int, t []float64, ldt int, work []float64) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtpqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctpqrt.f.
func Ctpqrt(m, n, l, nb int, a []complex64, lda int, b []complex64, ldb int, t []complex64, ldt int, work []complex64) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ctpqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztpqrt.f.
func Ztpqrt(m, n, l, nb int, a []complex128, lda int, b []complex128, ldb int, t []complex128, ldt int, work []complex128) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ztpqrt_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (C.lapack_int)(nb), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_work)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stpqrt2.f.
func Stpqrt2(m, n, l int, a []float32, lda int, b []float32, ldb int, t []float32, ldt int) bool {
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_stpqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtpqrt2.f.
func Dtpqrt2(m, n, l int, a []float64, lda int, b []float64, ldb int, t []float64, ldt int) bool {
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_dtpqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctpqrt2.f.
func Ctpqrt2(m, n, l int, a []complex64, lda int, b []complex64, ldb int, t []complex64, ldt int) bool {
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_ctpqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztpqrt2.f.
func Ztpqrt2(m, n, l int, a []complex128, lda int, b []complex128, ldb int, t []complex128, ldt int) bool {
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	return isZero(C.LAPACKE_ztpqrt2_work((C.int)(rowMajor), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(l), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/stprfb.f.
func Stprfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k, l int, v []float32, ldv int, t []float32, ldt int, a []float32, lda int, b []float32, ldb int, work []float32, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float32
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float32
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *float32
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float32
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float32
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_stprfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.float)(_v), (C.lapack_int)(ldv), (*C.float)(_t), (C.lapack_int)(ldt), (*C.float)(_a), (C.lapack_int)(lda), (*C.float)(_b), (C.lapack_int)(ldb), (*C.float)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/dtprfb.f.
func Dtprfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k, l int, v []float64, ldv int, t []float64, ldt int, a []float64, lda int, b []float64, ldb int, work []float64, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *float64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *float64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *float64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *float64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *float64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_dtprfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.double)(_v), (C.lapack_int)(ldv), (*C.double)(_t), (C.lapack_int)(ldt), (*C.double)(_a), (C.lapack_int)(lda), (*C.double)(_b), (C.lapack_int)(ldb), (*C.double)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ctprfb.f.
func Ctprfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k, l int, v []complex64, ldv int, t []complex64, ldt int, a []complex64, lda int, b []complex64, ldb int, work []complex64, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex64
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex64
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex64
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex64
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ctprfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.lapack_complex_float)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_float)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda), (*C.lapack_complex_float)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_float)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/ztprfb.f.
func Ztprfb(side blas.Side, trans blas.Transpose, direct, storev byte, m, n, k, l int, v []complex128, ldv int, t []complex128, ldt int, a []complex128, lda int, b []complex128, ldb int, work []complex128, ldwork int) bool {
	switch side {
	case blas.Left:
		side = 'L'
	case blas.Right:
		side = 'R'
	default:
		panic("lapack: bad side")
	}
	switch trans {
	case blas.NoTrans:
		trans = 'N'
	case blas.Trans:
		trans = 'T'
	case blas.ConjTrans:
		trans = 'C'
	default:
		panic("lapack: bad trans")
	}
	var _v *complex128
	if len(v) > 0 {
		_v = &v[0]
	}
	var _t *complex128
	if len(t) > 0 {
		_t = &t[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	var _b *complex128
	if len(b) > 0 {
		_b = &b[0]
	}
	var _work *complex128
	if len(work) > 0 {
		_work = &work[0]
	}
	return isZero(C.LAPACKE_ztprfb_work((C.int)(rowMajor), (C.char)(side), (C.char)(trans), (C.char)(direct), (C.char)(storev), (C.lapack_int)(m), (C.lapack_int)(n), (C.lapack_int)(k), (C.lapack_int)(l), (*C.lapack_complex_double)(_v), (C.lapack_int)(ldv), (*C.lapack_complex_double)(_t), (C.lapack_int)(ldt), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda), (*C.lapack_complex_double)(_b), (C.lapack_int)(ldb), (*C.lapack_complex_double)(_work), (C.lapack_int)(ldwork)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/csyr.f.
func Csyr(ul blas.Uplo, n int, alpha complex64, x []complex64, incx int, a []complex64, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _x *complex64
	if len(x) > 0 {
		_x = &x[0]
	}
	var _a *complex64
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_csyr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_complex_float)(alpha), (*C.lapack_complex_float)(_x), (C.lapack_int)(incx), (*C.lapack_complex_float)(_a), (C.lapack_int)(lda)))
}

// See http://www.netlib.org/cgi-bin/netlibfiles.txt?format=txt&filename=/lapack/lapack_routine/zsyr.f.
func Zsyr(ul blas.Uplo, n int, alpha complex128, x []complex128, incx int, a []complex128, lda int) bool {
	switch ul {
	case blas.Upper:
		ul = 'U'
	case blas.Lower:
		ul = 'L'
	default:
		panic("lapack: illegal triangle")
	}
	var _x *complex128
	if len(x) > 0 {
		_x = &x[0]
	}
	var _a *complex128
	if len(a) > 0 {
		_a = &a[0]
	}
	return isZero(C.LAPACKE_zsyr_work((C.int)(rowMajor), (C.char)(ul), (C.lapack_int)(n), (C.lapack_complex_double)(alpha), (*C.lapack_complex_double)(_x), (C.lapack_int)(incx), (*C.lapack_complex_double)(_a), (C.lapack_int)(lda)))
}

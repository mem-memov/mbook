// Package mbook is an in-memory storage that grows without copying stored data.
// Package name stands for "memory book".
// It's a set of pages. Pages may have varying size in bytes but all of them contain the same number of unsigned integers.
// Small unsigned integers require less space, so first pages are smaller in size than pages following after them.
// Pages are added automatically when new unsigned integers are written at a position in the book.
// Each page contains a column of lines, an unsigned integer per line. All lines in a column have the same size in bites.
package mbook

package helpers;

public enum Color {
	GREEN("\u001B[32m"),
	YELLOW("\u001B[33m"),
	RED("\u001B[31m"),
	RESET("\u001B[0m");

	final String val;

	Color(String val) {
		this.val = val;
	}

	@Override
	public String toString() {
		return this.val;
	}
}
